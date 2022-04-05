package mock_text_generator

import (
	"math/rand"
	"strings"
	"unicode"
)

type Strategy int

const (
	toLower = 0
	toUpper = 1
)

func ToMockText(boring string) string {
	return strings.Map(randomlyCased, boring)
}

func randomlyCased(toRandomize rune) rune {
	strategy := getCasingStrategy()
	if strategy == toLower {
		return unicode.ToLower(toRandomize)
	}
	return unicode.ToUpper(toRandomize)
}

func getCasingStrategy() Strategy {
	return Strategy(rand.Intn(2))
}
