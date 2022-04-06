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

	mux.Group(func(mux chi.Router) {
		mux.Use(app.session.Enable)
		mux.Get("/", app.home)
		mux.Get("/snippet/{id}", app.showSnippet)

		//Authentication
		mux.Get("/user/signup", app.signupUserForm)
		mux.Post("/user/signup", app.signupUser)
		mux.Get("/user/login", app.loginUserForm)
		mux.Post("/user/login", app.loginUser)
		// Require authentication

	})

	mux.Group(func(mux chi.Router) {
		mux.Use(app.session.Enable, app.requireAuthetication)
		mux.Get("/snippet/create", app.createSnippetForm)
		mux.Post("/snippet/create", app.createSnippet)
		mux.Post("/user/logout", app.logoutUser)
	})

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return app.logRequest(secureHeader(mux))
}
