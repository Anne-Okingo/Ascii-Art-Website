package sava

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"ascii-art-web/asciiArt"
)

func HandleasciiArt(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/template.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	// Check for only methodpost request.
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	// Return bad request status error.
	if text == "" || banner == "" {
		http.Error(w, " 400 Bad request", http.StatusBadRequest)
		return
	}

	os.MkdirAll("bannerfiles", 0o777)

	bannerMap, err := asciiArt.LoadBannerMap("bannerfiles/" + banner + ".txt")
	if err != nil {
		// return error 500
		errorMessage := fmt.Sprintf(" 404 Not Found \n%v", err)
		http.Error(w, errorMessage, http.StatusNotFound)
		return
	}
	art, err := asciiArt.PrintLineBanner(text, bannerMap)
	if err != nil {
		// return error 500// return error 500
		errorMessage := fmt.Sprintf(" 500 Internal Server Error \n%v", err)
		http.Error(w, errorMessage, http.StatusInternalServerError)
		return
	}

	data := struct {
		AsciiArt string
	}{
		AsciiArt: art,
	}

	t.ExecuteTemplate(w, "template.html", data)
}
