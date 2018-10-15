package main

import (
	"strings"
	"testing"
)

var words = []string{
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
	"22", "333", "4444", "55555", "666666", "7777777", "88888888", "999999999",
}

func TestGeneratePasswords(t *testing.T) {
	passwordsCount := 3
	wordsCount := 3

	passwords := GeneratePasswords(words, passwordsCount, wordsCount)
	if len(passwords) != int(passwordsCount) {
		t.Fatalf("expected %d passwords, got %d", passwordsCount, len(passwords))
	}

	for _, p := range passwords {
		tokens := strings.FieldsFunc(p, func(r rune) bool {
			return strings.ContainsRune(separators, r)
		})
		if len(tokens) != wordsCount {
			t.Fatalf("expected %d words, got %d (%s)", wordsCount, len(tokens), p)
		}
		for _, token := range tokens {
			if !findWord(words, token) {
				t.Fatalf("uexpected word %s", token)
			}
		}
	}
}

func TestGeneratePasswordsOnlyOne(t *testing.T) {
	passwordsCount := 1
	wordsCount := 1

	passwords := GeneratePasswords(words, passwordsCount, wordsCount)
	if len(passwords) != int(passwordsCount) {
		t.Fatalf("expected %d passwords, got %d", passwordsCount, len(passwords))
	}

	for _, p := range passwords {
		tokens := strings.FieldsFunc(p, func(r rune) bool {
			return strings.ContainsRune(separators, r)
		})
		if len(tokens) != wordsCount {
			t.Fatalf("expected %d words, got %d (%s)", wordsCount, len(tokens), p)
		}
		for _, token := range tokens {
			if !findWord(words, token) {
				t.Fatalf("unexpected word %s", token)
			}
		}
	}
}

func TestGeneratePasswordsFromEmptyList(t *testing.T) {
	words := []string{}
	passwordsCount := 3
	wordsCount := 3

	passwords := GeneratePasswords(words, passwordsCount, wordsCount)
	if len(passwords) != 0 {
		t.Fatalf("expected 0 passwords, got %d", len(passwords))
	}
}

func TestGeneratePasswordsZero(t *testing.T) {
	words := []string{}
	passwordsCount := 0
	wordsCount := 0

	passwords := GeneratePasswords(words, passwordsCount, wordsCount)
	if len(passwords) != 0 {
		t.Fatalf("expected 0 passwords, got %d", len(passwords))
	}
}

func findWord(words []string, word string) bool {
	for _, w := range words {
		if word == w {
			return true
		}
	}
	return false
}
