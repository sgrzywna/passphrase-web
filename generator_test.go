package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

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

// func TestLoadDicts(t *testing.T) {
// 	names := []string{"first.dict", "second.dict", "third.dict"}
// 	dictsFolder := getFolderWithDicts(t, words, names...)
// 	defer os.RemoveAll(dictsFolder)
// 	dicts, err := loadDicts(dictsFolder)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	for _, dict := range dicts {
// 		if !findWord(names, dict.fileName) {
// 			t.Errorf("unexpected dict name %s", dict.fileName)
// 		}
// 		for _, word := range dict.words {
// 			if !findWord(words, word) {
// 				t.Errorf("unexpected word %s", word)
// 			}
// 		}
// 	}
// }

func getFolderWithDicts(t *testing.T, words []string, names ...string) string {
	data := strings.Join(words, "\n")
	tmpDir := os.TempDir()
	for _, name := range names {
		if err := ioutil.WriteFile(filepath.Join(tmpDir, name), []byte(data), os.ModePerm); err != nil {
			t.Fatal(err)
		}
	}
	return tmpDir
}
