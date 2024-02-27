package pkg

import (
	"github.com/pemistahl/lingua-go"
)

type Detector struct{}

func NewDetector() *Detector {
	return &Detector{}
}

func (t *Detector) IsSupported(text string) bool {
	// supported languages by google gemini api
	langs := []lingua.Language{
		lingua.Arabic,
		lingua.Bengali,
		lingua.Bulgarian,
		lingua.Chinese,
		lingua.Croatian,
		lingua.Czech,
		lingua.Danish,
		lingua.Dutch,
		lingua.English,
		lingua.Estonian,
		lingua.Finnish,
		lingua.French,
		lingua.German,
		lingua.Greek,
		lingua.Hebrew,
		lingua.Hindi,
		lingua.Hungarian,
		lingua.Indonesian,
		lingua.Italian,
		lingua.Japanese,
		lingua.Korean,
		lingua.Latvian,
		lingua.Lithuanian,
		// lingua.Norwegian,   lingua does not support this
		lingua.Polish,
		lingua.Portuguese,
		lingua.Romanian,
		lingua.Russian,
		lingua.Serbian,
		lingua.Slovak,
		lingua.Slovene,
		lingua.Spanish,
		lingua.Swahili,
		lingua.Swedish,
		lingua.Thai,
		lingua.Turkish,
		lingua.Ukrainian,
		lingua.Vietnamese,
	}

	detector := lingua.
		NewLanguageDetectorBuilder().
		FromLanguages(langs...).
		Build()

	_, exists := detector.DetectLanguageOf(text)

	return exists

}
