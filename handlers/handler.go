package sava

import (
	"html/template"
	"net/http"
	"os"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Checks if the HTML template file exists
	if _, err := os.Stat("templates/template.html"); os.IsNotExist(err) {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	t, err := template.ParseFiles("templates/template.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	// define the data to return
	if r.URL.Path == "/" {
		t.ExecuteTemplate(w, "template.html", "nil")
	} else {
		http.NotFound(w, r)
	}
}
