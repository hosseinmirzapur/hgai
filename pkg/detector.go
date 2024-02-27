package pkg

import (
	"fmt"

	"github.com/pemistahl/lingua-go"
)

type Detector struct {
	text string
}

func NewDetector(text string) *Detector {
	return &Detector{
		text: text,
	}
}

func (t *Detector) DetectLanguage() (lingua.Language, error) {
	detector := lingua.NewLanguageDetectorBuilder().
		FromAllLanguages().
		Build()

	lang, exists := detector.DetectLanguageOf(t.text)

	if !exists {
		return 0, fmt.Errorf("input language not detected")
	}

	return lang, nil

}
