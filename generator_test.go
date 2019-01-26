package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var words = []string{
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
	"22", "333", "4444", "55555", "666666", "7777777", "88888888", "999999999",
}

func TestGetDictFiles(t *testing.T) {
	names := []string{"first.dict", "second.dict", "third.dict"}
	dictsFolder := getFolderWithDicts(t, words, names...)
	defer os.RemoveAll(dictsFolder)
	g, err := NewGenerator(dictsFolder)
	if err != nil {
		t.Fatal(err)
	}
	fileNames := g.GetDictFiles()
	for _, name := range fileNames {
		if !findWord(names, name) {
			t.Errorf("unexpected dict file name %s", name)
		}
	}
}

func TestGeneratePasswords(t *testing.T) {
	dictsFolder := getDictsFolder(t)
	defer os.RemoveAll(dictsFolder)

	g, err := NewGenerator(dictsFolder)
	if err != nil {
		t.Fatal(err)
	}

	passwordsCount := 3
	wordsCount := 3

	passwords := g.getPasswords(words, passwordsCount, wordsCount)
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
	dictsFolder := getDictsFolder(t)
	defer os.RemoveAll(dictsFolder)

	g, err := NewGenerator(dictsFolder)
	if err != nil {
		t.Fatal(err)
	}

	passwordsCount := 1
	wordsCount := 1

	passwords := g.getPasswords(words, passwordsCount, wordsCount)
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
	dictsFolder := getDictsFolder(t)
	defer os.RemoveAll(dictsFolder)

	g, err := NewGenerator(dictsFolder)
	if err != nil {
		t.Fatal(err)
	}

	words := []string{}
	passwordsCount := 3
	wordsCount := 3

	passwords := g.getPasswords(words, passwordsCount, wordsCount)
	if len(passwords) != 0 {
		t.Fatalf("expected 0 passwords, got %d", len(passwords))
	}
}

func TestGeneratePasswordsZero(t *testing.T) {
	dictsFolder := getDictsFolder(t)
	defer os.RemoveAll(dictsFolder)

	g, err := NewGenerator(dictsFolder)
	if err != nil {
		t.Fatal(err)
	}

	words := []string{}
	passwordsCount := 0
	wordsCount := 0

	passwords := g.getPasswords(words, passwordsCount, wordsCount)
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

func getFolderWithDicts(t *testing.T, words []string, names ...string) string {
	data := strings.Join(words, "\n")
	tmpDir := getDictsFolder(t)
	for _, name := range names {
		if err := ioutil.WriteFile(filepath.Join(tmpDir, name), []byte(data), os.ModePerm); err != nil {
			t.Fatal(err)
		}
	}
	return tmpDir
}

func getDictsFolder(t *testing.T) string {
	tmpDir := filepath.Join(os.TempDir(), "dicts")
	if err := os.MkdirAll(tmpDir, os.ModePerm); err != nil {
		t.Fatal(err)
	}
	return tmpDir
}
