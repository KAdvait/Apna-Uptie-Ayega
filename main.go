package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"time"
)

func statusPageHandler(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("templates", "status.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, statuses)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
func main() {
	go func() {
		for range time.Tick(10 * time.Second) {
			checkAllEndpoints()
		}
	}()

	http.HandleFunc("/", statusPageHandler)
	http.ListenAndServe(":8090", nil)
}
