package main

import (
	"bufio"
	"os"
	"strings"
)

const (
	maxPasswordsCount int    = 128
	maxWordsCount     int    = 16
	separators        string = " +-=:.,!#%&*_"
)

// GeneratePasswordsFromFile returns list of random passwords generated from dictionary stored in file.
func GeneratePasswordsFromFile(dictionary string, passwordsCount int, wordsCount int) ([]string, error) {
	if passwordsCount < 1 {
		passwordsCount = 1
	} else if passwordsCount > maxPasswordsCount {
		passwordsCount = maxPasswordsCount
	}

	if wordsCount < 1 {
		wordsCount = 1
	} else if wordsCount > maxWordsCount {
		wordsCount = maxWordsCount
	}

	d, err := LoadDictionary(dictionary)
	if err != nil {
		return nil, err
	}

	passwords := GeneratePasswords(d, passwordsCount, wordsCount)
	return passwords, nil
}

// GeneratePasswords returns list of random passwords generated from dictionary stored in memory.
func GeneratePasswords(dictionary []string, passwordsCount int, wordsCount int) []string {
	if len(dictionary) == 0 {
		return []string{}
	}

	if passwordsCount < 1 {
		passwordsCount = 1
	} else if passwordsCount > maxPasswordsCount {
		passwordsCount = maxPasswordsCount
	}

	if wordsCount < 1 {
		wordsCount = 1
	} else if wordsCount > maxWordsCount {
		wordsCount = maxWordsCount
	}

	var passwords []string
	maxNumber := uint32(len(dictionary) - 1)

	r := NewRandIndex()
	for i := 0; i < passwordsCount; i++ {
		var words []string
		for j := 0; j < wordsCount; j++ {
			ndx := r.RandInt32(maxNumber)
			words = append(words, dictionary[ndx])
		}
		sep := separators[r.RandInt32(uint32(len(separators))-1)]
		passwords = append(passwords, strings.Join(words, string(sep)))
	}

	return passwords
}

// LoadDictionary loads dictionary words into memory.
func LoadDictionary(dictionary string) ([]string, error) {
	file, err := os.Open(dictionary)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
