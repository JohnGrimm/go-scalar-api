package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/JohnGrimm/go-scalar-api"
)

// @title           Simple API
// @version         1.0
// @description     Exemple use of scalar beautfull api
// @termsOfService  http://swagger.io/terms/

// @contact.name   Marcelo Petrucio
// contact.url    https://marcelopetrucio.dev
// @contact.email  marcelo.petrucio43@gmail.com

// @BasePath  /
func main() {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/", create)

	router.Get("/docs", func(w http.ResponseWriter, r *http.Request) {
		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL: "./docs/swagger.json",
			CustomOptions: scalar.CustomOptions{
				PageTitle: "Simple API",
			},
			DarkMode: true,
		})

		if err != nil {
			fmt.Printf("%v", err)
		}

		fmt.Fprintln(w, htmlContent)
	})

	fmt.Printf("Starting web server on port :8000")
	http.ListenAndServe(":8000", router)
}
