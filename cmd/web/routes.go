package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

//
//func (app *application) routes() http.Handler {
//	mux := http.NewServeMux()
//
//	mux.HandleFunc("/", app.home)
//	mux.HandleFunc("/snippet", app.showSnippet)
//	mux.HandleFunc("/snippet/create", app.createSnippet)
//
//	fileServer := http.FileServer(http.Dir("./ui/static"))
//	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
//
//	return app.logRequest(secureHeader(mux))
//}

func (app *application) routes() http.Handler {
	mux := chi.NewMux()

	mux.Get("/", app.home)
	mux.Get("/snippet/{id}", app.showSnippet)
	mux.Post("/snippet/create", app.createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return app.logRequest(secureHeader(mux))
}
