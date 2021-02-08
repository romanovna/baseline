package main

import (
	driver "crud-api/driver/mysql"
	handler "crud-api/handler/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
)

func main() {
	dbName := "posts"
	dbPass := "password123"
	dbHost := "localhost"
	dbPort := "3306"
	dbUser := "root"

	conn, err := driver.ConnectSQL(dbHost, dbPort, dbUser, dbPass, dbName)
	if err != nil {
		panic(err.Error())
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	pHandler := handler.NewPostHandler(conn)
	r.Route("/", func(rt chi.Router) {
		rt.Mount("/posts", postRouter(pHandler))
	})
	log.Println("Server listen at :8080")
	http.ListenAndServe(":8080", r)
}

func postRouter(pHandler *handler.Post) http.Handler {
	r := chi.NewRouter()
	r.Get("/", pHandler.Fetch)
	r.Get("/{id:[0-9]+}", pHandler.GetByID)
	r.Post("/", pHandler.Create)
	r.Put("/{id:[0-9]+}", pHandler.Update)
	r.Delete("/{id:[0-9]+}", pHandler.Delete)
	r.Post("/init", pHandler.Init)
	return r
}
