package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	dictFileSuffix           = ".dict"
	maxPasswordsCount int    = 128
	maxWordsCount     int    = 16
	separators        string = " +-=:.,!#%&*_"
)

// Generator represents password generator.
type Generator struct {
	dicts     map[string][]string
	randIndex *RandIndex
}

// NewGenerator returns initialized Generator object.
func NewGenerator(dictsDir string) (*Generator, error) {
	generator := Generator{
		dicts:     make(map[string][]string),
		randIndex: NewRandIndex(),
	}
	if err := generator.loadDicts(dictsDir); err != nil {
		return nil, err
	}
	return &generator, nil
}

// GetDictFiles returns list of loaded dictionary files.
func (g *Generator) GetDictFiles() []string {
	var names []string
	for name := range g.dicts {
		names = append(names, name)
	}
	return names
}

// GeneratePasswords returns list of random passwords.
func (g *Generator) GeneratePasswords(dict string, passwords, words int) []string {
	dictionary, ok := g.dicts[dict]
	if !ok {
		return []string{}
	}
	return g.getPasswords(dictionary, passwords, words)
}

// loadDicts loads dictionary files into memory.
func (g *Generator) loadDicts(dictsDir string) error {
	files, err := ioutil.ReadDir(dictsDir)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), dictFileSuffix) {
			continue
		}
		words, err := g.loadDictionary(filepath.Join(dictsDir, file.Name()))
		if err != nil {
			return err
		}
		g.dicts[strings.ToLower(file.Name())] = words
	}
	return nil
}

// GeneratePasswordsFromFile returns list of random passwords generated from dictionary stored in file.
func (g *Generator) generatePasswordsFromFile(dictionary string, passwordsCount int, wordsCount int) ([]string, error) {
	d, err := g.loadDictionary(dictionary)
	if err != nil {
		return nil, err
	}
	passwords := g.getPasswords(d, passwordsCount, wordsCount)
	return passwords, nil
}

// getPasswords returns list of random passwords generated from dictionary stored in memory.
func (g *Generator) getPasswords(dictionary []string, passwordsCount int, wordsCount int) []string {
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

	for i := 0; i < passwordsCount; i++ {
		var words []string
		for j := 0; j < wordsCount; j++ {
			ndx := g.randIndex.RandInt32(maxNumber)
			words = append(words, dictionary[ndx])
		}
		sep := separators[g.randIndex.RandInt32(uint32(len(separators))-1)]
		passwords = append(passwords, strings.Join(words, string(sep)))
	}

	return passwords
}

// loadDictionary loads dictionary words into memory.
func (g *Generator) loadDictionary(dictionary string) ([]string, error) {
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
