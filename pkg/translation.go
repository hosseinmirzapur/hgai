package pkg

import (
	"fmt"

	"github.com/bregydoc/gtranslate"
	"github.com/pemistahl/lingua-go"
)

type Translation struct {
	Text     string
	Language string
}

func NewTranslation(text string) *Translation {
	return &Translation{
		Text: text,
	}
}

func (t *Translation) Translate(from, to string) (string, error) {
	return gtranslate.TranslateWithParams(t.Text, gtranslate.TranslationParams{
		From: from,
		To:   to,
	})

}

func (t *Translation) DetectLanguage() (string, error) {
	detector := lingua.NewLanguageDetectorBuilder().
		FromAllLanguages().
		Build()

	lang, exists := detector.DetectLanguageOf(t.Text)

	if !exists {
		return "", fmt.Errorf("input language not detected")
	}

	t.Language = lang.IsoCode639_1().String()

	return t.Language, nil

}
