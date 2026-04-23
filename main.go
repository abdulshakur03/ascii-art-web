package main

import (
	asciiart "ascii-art-web/ascii-art"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	temp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	temp.Execute(w, nil)

}

func asciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	bannerName := r.FormValue("banner")

	if text == "" {
		http.Error(w, "Text cannot be empty", http.StatusBadRequest)
		return
	}

	if bannerName != "standard" && bannerName != "shadow" && bannerName != "thinkertoy" {
		http.Error(w, "Invalid banner selection", http.StatusBadRequest)
		return
	}

	banner, err := asciiart.ReadBannerFile("banners/" + bannerName + ".txt")
	if err != nil {
		http.Error(w, "Banner not found", http.StatusNotFound)
		return
	}

	lines := asciiart.SplitInputLines(text)

	ascii := asciiart.BuildAscii(lines, banner)

	temp, err := template.ParseFiles("templates/result.htm")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Result string
	}{
		Result: ascii,
	}

	if err := temp.Execute(w, data); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiHandler)
	fmt.Println("listening on Port: 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
