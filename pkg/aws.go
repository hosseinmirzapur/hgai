package pkg

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/translate"
)

type AWS struct {
	client *translate.Translate
}

func NewAWS() *AWS {
	// establish a new session
	sess := session.Must(session.NewSession())

	// AWS client for translation
	client := translate.New(sess)

	return &AWS{
		client: client,
	}
}

func (a *AWS) Translate(text string) (string, error) {
	result, err := a.client.Text(&translate.TextInput{
		SourceLanguageCode: aws.String("auto"),
		TargetLanguageCode: aws.String("en"),
		Text:               aws.String(text),
	})

	if err != nil {
		return "", err
	}

	return *result.TranslatedText, nil
}
