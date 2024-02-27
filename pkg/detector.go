package pkg

import (
	"fmt"

	"github.com/pemistahl/lingua-go"
)

type Detector struct{}

func NewDetector() *Detector {
	return &Detector{}
}

func (t *Detector) DetectLanguage(text string) (lingua.Language, error) {
	detector := lingua.NewLanguageDetectorBuilder().
		FromAllLanguages().
		Build()

	lang, exists := detector.DetectLanguageOf(text)

	if !exists {
		return 0, fmt.Errorf("input language not detected")
	}

	return lang, nil

}
