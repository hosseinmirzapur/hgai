package pkg

import (
	"fmt"
	"strings"

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

func (t *Detector) DetectLanguage() (string, error) {
	detector := lingua.NewLanguageDetectorBuilder().
		FromAllLanguages().
		Build()

	lang, exists := detector.DetectLanguageOf(t.text)

	if !exists {
		return "", fmt.Errorf("input language not detected")
	}

	return strings.ToLower(lang.IsoCode639_1().String()), nil

}
