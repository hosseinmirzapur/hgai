package pkg

import "github.com/bregydoc/gtranslate"

type Translate struct{}

func NewTranslation() *Translate {
	return &Translate{}
}

func (t *Translate) ToEnglish(text string) (string, error) {
	return gtranslate.TranslateWithParams(text, gtranslate.TranslationParams{
		From: "auto",
		To:   "en",
	})
}
