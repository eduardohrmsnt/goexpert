package main

import (
	"log"
	"net/http"

	"github.com/eduardohrmsnt/go-expert-api/configs"
	"github.com/eduardohrmsnt/go-expert-api/internal/entity"
	"github.com/eduardohrmsnt/go-expert-api/internal/infra/database"
	"github.com/eduardohrmsnt/go-expert-api/internal/infra/webserver/handlers"
	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})

	mux := chi.NewRouter()

	productHandler := handlers.NewProductHandler(database.NewProduct(db))
	userHandler := handlers.NewUserHandler(database.NewUser(db))

	mux.Use(middleware.Logger)

	mux.Use(middleware.WithValue("jwt", config.TokenAuth))
	mux.Use(middleware.WithValue("expiresIn", config.JwtExperesIn))

	mux.Route("/products", func(mux chi.Router) {
		mux.Use(jwtauth.Verifier(config.TokenAuth))
		mux.Use(jwtauth.Authenticator)
		mux.Post("/", productHandler.CreateProduct)
		mux.Get("/", productHandler.GetProducts)
		mux.Get("/{id}", productHandler.GetProduct)
		mux.Put("/{id}", productHandler.UpdateProduct)
		mux.Delete("/{id}", productHandler.DeleteProduct)
	})

	mux.Post("/users", userHandler.Create)
	mux.Post("/users/generate_token", userHandler.GenerateToken)

	http.ListenAndServe(":8080", mux)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("Request: %s %s", req.Method, req.URL.Path)
		next.ServeHTTP(w, req)
	})
}
