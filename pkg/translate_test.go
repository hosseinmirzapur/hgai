package pkg

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslation(t *testing.T) {
	translator := NewTranslation()

	translated, err := translator.ToEnglish("سلام حالت چطوره؟ من اسمم حسین هست اسم تو چیه؟")
	if err != nil {
		log.Fatal("unable to translate the input text")
	}
	log.Println(translated)

	assert.Contains(t, translated, "name")

}
