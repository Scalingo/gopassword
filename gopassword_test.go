package gopassword

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	t.Run("When we want to generate a password", func(t *testing.T) {
		t.Run("By default, it must be 24 characters", func(t *testing.T) {
			assert.Len(t, Generate(), 24)
		})

		t.Run("With an argument, the generated password must have its length", func(t *testing.T) {
			assert.Len(t, Generate(10), 10)
			assert.Len(t, Generate(42), 42)
			assert.Len(t, Generate(999), 999)
		})

		t.Run("With several arguments, only the first must be considered", func(t *testing.T) {
			assert.Len(t, Generate(10, 20, 30), 10)
		})

		t.Run("It must contain only alphanumeric or underscore characters", func(t *testing.T) {
			allowedCharacters := regexp.MustCompile("^[a-zA-Z0-9_]+$")

			// Try various times to ensure the result is not casual
			for _ = range 1000 {
				passwd := Generate(99)
				assert.True(t, allowedCharacters.MatchString(passwd))
			}
		})
	})

	t.Run("Given a generated password", func(t *testing.T) {
		passwd := Generate(20)
		t.Run("The character frequency must be low", func(t *testing.T) {
			fm := frequencyMap(passwd)
			maxFreq := max(fm)
			assert.LessOrEqual(t, maxFreq, 3)
		})
	})
}

func frequencyMap(p string) map[rune]int {
	charMap := make(map[rune]int)
	for _, r := range p {
		charMap[r] += 1
	}
	return charMap
}

func max(m map[rune]int) int {
	res := 0
	for _, v := range m {
		if v > res {
			res = v
		}
	}
	return res
}
