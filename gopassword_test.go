package gopassword

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	t.Run("When we want to generate a password", func(t *testing.T) {
		t.Run("By default, it should be 20 characters", func(t *testing.T) {
			assert.Len(t, Generate(), 20)
		})

		t.Run("With an argument, the generated password should have its length", func(t *testing.T) {
			assert.Len(t, Generate(42), 42)
		})

		t.Run("With several arguments, only the first should be considered", func(t *testing.T) {
			assert.Len(t, Generate(10, 20, 30), 10)
		})
	})

	t.Run("Given a generated password", func(t *testing.T) {
		passwd := Generate(20)
		t.Run("The character frequency should be low", func(t *testing.T) {
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
