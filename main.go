package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func main() {
	port := flag.Int("port", 8080, "listening port")
	dir := flag.String("dir", ".", "static files directory")
	dictDir := flag.String("dicts", ".", "dictionary files directory")
	flag.Parse()
	log.Printf("Listening @ :%d...", *port)
	log.Printf("Files directory: %s", *dir)

	staticDir := filepath.Join(*dir, "static")
	log.Printf("Static files directory: %s", staticDir)

	t, err := template.ParseFiles(filepath.Join(*dir, "index.html"))
	if err != nil {
		log.Fatal(err)
	}

	generator, err := NewGenerator(*dictDir)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))
	http.HandleFunc("/passwords.json", limit(getPasswords(generator)))
	http.HandleFunc("/", getIndexHandler(generator, t))

	initLimiter()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

func getIndexHandler(generator *Generator, t *template.Template) http.HandlerFunc {
	type selectData struct {
		Name  string
		Value string
	}

	var dictNames []selectData

	files := generator.GetDictFiles()
	sort.Strings(files)

	for _, file := range files {
		name := strings.Title(strings.TrimSuffix(file, filepath.Ext(file)))
		dictNames = append(dictNames, selectData{name, file})
	}

	return func(w http.ResponseWriter, r *http.Request) {
		err := t.Execute(w, dictNames)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

func getPasswords(generator *Generator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		values := r.URL.Query()
		dict := getStringValue(values, "d", "")
		passwords := getIntValue(values, "p", 7)
		words := getIntValue(values, "w", 3)

		res := generator.GeneratePasswords(dict, passwords, words)
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

func getIntValue(values url.Values, key string, defaultValue int) int {
	val := values.Get(key)
	if val != "" {
		res, err := strconv.Atoi(val)
		if err == nil {
			return res
		}
	}
	return defaultValue
}

func getStringValue(values url.Values, key, defaultValue string) string {
	val := values.Get(key)
	if val != "" {
		return val
	}
	return defaultValue
}
