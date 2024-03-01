package pkg

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/generative-ai-go/genai"
)

func handleDefaultCommand(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Invalid command. Send /help to get bot help info")
	sendMessage(bot, msg)
}

func handleStartCommand(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	session := NewSession()
	dynamo := NewDynamoDB(session)
	chatID := update.Message.Chat.ID

	msg, err := RegisterNewUser(dynamo, update.Message.From.ID)
	if err != nil {
		handleErrorViaBot(bot, chatID, err)
		return
	}
	sendMessage(bot, tgbotapi.NewMessage(chatID, msg))

	botUser := update.Message.From
	startText := fmt.Sprintf("Hi %s!, Welcome to Smartinex Bot! Send /help to get help info", botUser.FirstName)

	sendMessage(bot, tgbotapi.NewMessage(chatID, startText))
}

func handleClearCommand(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	chatID := update.Message.Chat.ID
	textSessionID := generateSessionID(chatID, TextModel)

	info := "no chat session found, just send text or image"
	if ok := clearChatSession(textSessionID); ok {
		info = `Chat session cleared.`
	}

	msg := tgbotapi.NewMessage(chatID, info)
	sendMessage(bot, msg)
}

func handleHelpCommand(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	helpInfo := `Commands: 
    /clear - Clear chat session
    /help - Get help info
Just send text or image to get response`
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, helpInfo)
	sendMessage(bot, msg)
}

func handleTextMessage(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	chatID := update.Message.Chat.ID
	textPrompt := update.Message.Text

	initMsgID, errFlag := instantReply(update, bot, chatID)
	if errFlag {
		return
	}

	session := NewSession()

	comprehend := NewComprehend(session)
	inputLang, err := DetectLanguage(comprehend, textPrompt)
	if err != nil {
		handleErrorViaBot(bot, chatID, err)
		return
	}

	translate := NewTranslate(session)
	translatedPrompt, err := TranslateTo(
		translate,
		textPrompt,
		inputLang,
		"en",
	)

	if err != nil {
		handleErrorViaBot(bot, chatID, err)
		return
	}

	textPrompt = translatedPrompt

	a := &AWS{
		sess:      session,
		trans:     translate,
		compr:     comprehend,
		inputLang: inputLang,
	}

	generateResponse(a, bot, chatID, initMsgID, TextModel, genai.Text(textPrompt))

}

func handlePhotoMessage(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	chatID := update.Message.Chat.ID

	initMsgID, errFlag := instantReply(update, bot, chatID)
	if errFlag {
		return
	}

	var prompts []genai.Part
	a, errFlag := handlePhotoPrompts(update, bot, &prompts)
	if errFlag {
		return
	}

	generateResponse(a, bot, chatID, initMsgID, VisionModel, prompts...)
}

func instantReply(update tgbotapi.Update, bot *tgbotapi.BotAPI, chatID int64) (int, bool) {
	msg := tgbotapi.NewMessage(chatID, "Thinking...")
	msg.ReplyToMessageID = update.Message.MessageID
	initMsg, err := bot.Send(msg)
	if err != nil {
		log.Printf("Error sending message: %v\n", err)
		return 0, true
	}
	// Simulate typing action.
	_, _ = bot.Send(tgbotapi.NewChatAction(chatID, tgbotapi.ChatTyping))

	return initMsg.MessageID, false
}

func handlePhotoPrompts(update tgbotapi.Update, bot *tgbotapi.BotAPI, prompts *[]genai.Part) (*AWS, bool) {
	photo := update.Message.Photo[len(update.Message.Photo)-1]

	photoURL, err := getURL(bot, photo.FileID)
	if err != nil {
		return nil, true
	}
	imgData, err := getImageData(photoURL)
	if err != nil {
		return nil, true
	}
	imgType := getImageType(imgData)
	*prompts = append(*prompts, genai.ImageData(imgType, imgData))

	textPrompts := update.Message.Caption
	if textPrompts == "" {
		textPrompts = "Analyse the image and Describe it in English"
	}

	session := NewSession()

	comprehend := NewComprehend(session)
	inputLang, err := DetectLanguage(comprehend, textPrompts)
	if err != nil {
		handleErrorViaBot(bot, update.Message.Chat.ID, err)
		return nil, true
	}

	translate := NewTranslate(session)
	translated, err := TranslateTo(translate, textPrompts, inputLang, "en")
	if err != nil {
		handleErrorViaBot(bot, update.Message.Chat.ID, err)
		return nil, true
	}
	textPrompts = translated

	a := &AWS{
		sess:      session,
		trans:     translate,
		compr:     comprehend,
		inputLang: inputLang,
	}

	*prompts = append(*prompts, genai.Text(textPrompts))
	return a, false
}

func getURL(bot *tgbotapi.BotAPI, fileID string) (string, error) {
	url, err := bot.GetFileDirectURL(fileID)
	if err != nil {
		log.Printf("Error getting img URL: %v\n", err)
		return "", err
	}
	return url, nil
}

func getImageData(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Printf("Error getting image response: %v\n", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error closing image response: %v\n", err)
		}
	}(res.Body)

	imgData, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading image data: %v", err)
		return nil, err
	}

	return imgData, nil
}

func getImageType(data []byte) string {
	mimeType := http.DetectContentType(data)
	imageType := "jpeg"
	if strings.HasPrefix(mimeType, "image/") {
		imageType = strings.Split(mimeType, "/")[1]
	}
	return imageType
}

func sendMessage(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	_, err := bot.Send(msg)
	if err != nil {
		handleErrorViaBot(bot, msg.ChatID, fmt.Errorf("typing status error"))
		return
	}
}

func generateResponse(a *AWS, bot *tgbotapi.BotAPI, chatID int64, initMsgID int, modelName string, parts ...genai.Part) {
	response := getModelResponse(chatID, modelName, parts)
	translated, err := TranslateTo(a.trans, response, "en", a.inputLang)
	if err != nil {
		handleErrorViaBot(bot, chatID, fmt.Errorf("session expired! send /clear to reset your session"))
		return
	}

	// Send the response back to the user.
	edit := tgbotapi.NewEditMessageText(chatID, initMsgID, translated)
	edit.ParseMode = tgbotapi.ModeMarkdown
	edit.DisableWebPagePreview = true

	if _, err := bot.Send(edit); err != nil {
		edit.ParseMode = ""

		if _, err = bot.Send(edit); err != nil {
			handleErrorViaBot(bot, chatID, fmt.Errorf("unable to send response"))
			return
		}
	}
}
