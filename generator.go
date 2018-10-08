package main

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

const (
	dictFileSuffix = ".dict"
)

// Generator represents password generator.
type Generator struct {
	dicts map[string][]string
}

// NewGenerator returns initialized Generator object.
func NewGenerator(dictsDir string) (*Generator, error) {
	generator := Generator{
		dicts: make(map[string][]string),
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
	return GeneratePasswords(dictionary, passwords, words)
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
		words, err := LoadDictionary(filepath.Join(dictsDir, file.Name()))
		if err != nil {
			return err
		}
		g.dicts[strings.ToLower(file.Name())] = words
	}
	return nil
}
