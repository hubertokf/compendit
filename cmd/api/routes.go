package main

import (
	"compendit/cmd/api/controllers"
	"compendit/cmd/api/helpers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// RegisterRoutes add all routing list here automatically get main router
func RegisterRoutes(router *chi.Mux) {

	// start router
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		helpers.SuccessResponse(w, "alive ok")
	})
	//Add All route
	router.Group(func(r chi.Router) {
		r.Post("/test/", controllers.CreateExample)
		r.Get("/test/", controllers.GetData)
	})
}

// end router
