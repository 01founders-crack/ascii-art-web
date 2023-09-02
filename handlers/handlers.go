// handlers.go

package handlers

import (
	ascii_art "ascii-art-web/ascii-art"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// renderTemplateWithData renders a template with given data.
func renderTemplateWithData(w http.ResponseWriter, tmplName string, data interface{}) {
	lp := filepath.Join("templates", "layout.html")
	fp := filepath.Join("templates", tmplName)

	info, err := os.Stat(fp)
	if err != nil || info.IsDir() {
		renderErrorTemplate(w, http.StatusNotFound, "errors/404.html")
		return
	}

	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		renderErrorTemplate(w, http.StatusInternalServerError, "errors/500.html")
		log.Print(err.Error())
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		renderErrorTemplate(w, http.StatusInternalServerError, "errors/500.html")
		log.Print(err.Error())
	}
}

func ServeTemplate(w http.ResponseWriter, r *http.Request) {
	// If the request is for "/index.html", redirect to "/"
	if r.URL.Path == "/index.html" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	// Define the layout path
	lp := filepath.Join("templates", "layout.html")

	// Determine the template file path based on the URL path
	var fp string
	if r.URL.Path == "/" {
		fp = filepath.Join("templates", "index.html")
	} else {
		fp = filepath.Join("templates", filepath.Clean(r.URL.Path))
	}

	// Check if the template file exists
	info, err := os.Stat(fp)
	if err != nil || info.IsDir() {
		renderErrorTemplate(w, http.StatusNotFound, "errors/404.html")
		return
	}

	// Parse and execute the template
	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		renderErrorTemplate(w, http.StatusInternalServerError, "errors/500.html")
		log.Print(err.Error())
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		renderErrorTemplate(w, http.StatusInternalServerError, "errors/500.html")
		log.Print(err.Error())
	}
}

func HandleAsciiArt(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest) //400 error
		return
	}

	artStyle := r.FormValue("artstyle")
	userText := r.FormValue("text")
	artStylePath := "ascii-art/artstyles/" + artStyle + ".txt"
	asciiArtResult := ascii_art.AsciiArt(userText, artStylePath)
	data := struct {
		ASCIIArtResult string
	}{
		ASCIIArtResult: asciiArtResult,
	}
	tmplName := "index.html"
	renderTemplateWithData(w, tmplName, data)
}

func renderErrorTemplate(w http.ResponseWriter, statusCode int, templatePath string) {
	w.WriteHeader(statusCode)
	renderTemplateWithData(w, templatePath, nil)
}
