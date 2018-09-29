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
	"strconv"
	"strings"
)

func main() {
	port := flag.Int("port", 8080, "listening port")
	dir := flag.String("dir", ".", "static files directory")
	flag.Parse()
	log.Printf("Listening @ :%d...", *port)
	log.Printf("Files directory: %s", *dir)

	staticDir := filepath.Join(*dir, "static")
	log.Printf("Static files directory: %s", staticDir)

	t, err := template.ParseFiles(filepath.Join(*dir, "index.html"))
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))
	http.HandleFunc("/passwords.json", getPasswords())
	http.HandleFunc("/", getIndexHandler(t))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

func getIndexHandler(t *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := t.Execute(w, struct{}{})
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

func getPasswords() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		values := r.URL.Query()
		words := getIntValue(values, "w", 3)
		passwords := getIntValue(values, "p", 7)

		var res []string

		for i := 0; i < passwords; i++ {
			var w []string
			for j := 0; j < words; j++ {
				w = append(w, "word")
			}
			res = append(res, strings.Join(w, "-"))
		}

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
