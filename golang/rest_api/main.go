package main

import (
	"crud-api/db"
	"crud-api/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	openApiMid "github.com/go-openapi/runtime/middleware"
	"log"
	"net/http"
)

func main() {
	dbName := "posts"
	dbPass := "password123"
	dbHost := "localhost"
	dbPort := "3306"
	dbUser := "root"

	conn, err := db.ConnectSQL(dbHost, dbPort, dbUser, dbPass, dbName)
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(conn)
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	opts := openApiMid.RedocOpts{SpecURL: "/swagger.yaml"}
	swaggerDocs := openApiMid.Redoc(opts, nil)

	pHandler := handler.NewHandler(conn)
	r.Route("/", func(rt chi.Router) {
		rt.Mount("/docs", swaggerDocs)
		rt.Mount("/swagger.yaml", http.FileServer(http.Dir("./")))
		rt.Mount("/posts", postRouter(pHandler))
	})
	log.Println("Server listen at :8080")
	http.ListenAndServe(":8080", r)
}

func postRouter(pHandler *handler.Handler) http.Handler {
	r := chi.NewRouter()
	r.Get("/", pHandler.Fetch)
	r.Get("/{id:[0-9]+}", pHandler.GetByID)
	r.Post("/", pHandler.Create)
	r.Put("/{id:[0-9]+}", pHandler.Update)
	r.Delete("/{id:[0-9]+}", pHandler.Delete)
	return r
}
