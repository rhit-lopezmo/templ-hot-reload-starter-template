package main

import (
	"net/http"
	"path/filepath"
	"tmpl-first-time/components"
	"tmpl-first-time/utils"

	"github.com/a-h/templ"
)

var log *utils.Logger

func main() {
	utils.InitLoggers()
	log = utils.MainLogger

	firstComp := components.FirstComponent("manuel lopez")

	http.Handle("/", templ.Handler(firstComp))

	staticFiles := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", staticFiles))

	// allows adjacent .js, .css, ..etc files next to their components
	http.Handle("/static/components/", http.StripPrefix("/static/components/",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ext := filepath.Ext(r.URL.Path)
			if ext == ".go" || ext == ".templ" {
				http.NotFound(w, r)
				return
			}
			http.FileServer(http.Dir("./components")).ServeHTTP(w, r)
		}),
	))

	log.Println("Listening on 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln("ERROR: could not start server", err)
	}
}
