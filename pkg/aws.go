package pkg

import (
	"cmp"
	"slices"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/comprehend"
	"github.com/aws/aws-sdk-go/service/translate"
)

type AWS struct {
	sess  *session.Session
	trans *translate.Translate
	compr *comprehend.Comprehend

	inputLang string
}

// establish a new session for aws services
func NewSession() *session.Session {
	sess := session.Must(session.NewSession())

	return sess
}

// return a new comprehend service instance
func NewComprehend(sess *session.Session) *comprehend.Comprehend {
	svc := comprehend.New(sess)

	return svc
}

// return a new translate service instance
func NewTranslate(sess *session.Session) *translate.Translate {
	svc := translate.New(sess)

	return svc
}

// detect the language of the input text
func DetectLanguage(cmpr *comprehend.Comprehend, text string) (string, error) {
	input := &comprehend.DetectDominantLanguageInput{
		Text: &text,
	}

	output, err := cmpr.DetectDominantLanguage(input)
	if err != nil {
		return "", err
	}

	dominantLang := slices.MaxFunc(output.Languages, func(a, b *comprehend.DominantLanguage) int {
		return cmp.Compare(*a.Score, *b.Score)
	})

	return *dominantLang.LanguageCode, nil

}

// translate text from any language to any language
// that aws translate service supports
func TranslateTo(trans *translate.Translate, text, from, to string) (string, error) {
	output, err := trans.Text(&translate.TextInput{
		SourceLanguageCode: &from,
		TargetLanguageCode: &to,
		Text:               &text,
	})

	if err != nil {
		return "", err
	}

	return *output.TranslatedText, nil
}
