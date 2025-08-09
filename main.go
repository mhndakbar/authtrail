package main

import (
	"database/sql"
	"embed"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mhndakbar/authtrails/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

//go:embed static/*
var staticFiles embed.FS

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Printf("PORT is not set in the .env file")
	}

	apiCfg := apiConfig{}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Println("DATABASE_URL is not set in the .env file")
		log.Println("Running without database")
	} else {
		db, err := sql.Open("postgres", dbURL)
		if err != nil {
			log.Fatal(err)
		}

		dbQueries := database.New(db)
		apiCfg.DB = dbQueries
		log.Println("Connected to database!")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		f, err := staticFiles.Open("static/index.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer f.Close()
		_, err = io.Copy(w, f)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})

	v1Router := chi.NewRouter()

	if apiCfg.DB != nil {
		v1Router.Post("/user", apiCfg.handlerCreateUser)
		v1Router.Get("/user", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
		v1Router.Post("/user/login", apiCfg.middlewareAuth(apiCfg.handlerLogin))
		v1Router.Get("/{userID}/authtrails", apiCfg.handlerGetAuthTrail)
	}

	router.Mount("/v1", v1Router)
	srv := &http.Server{
		Addr:        ":" + port,
		Handler:     router,
		ReadTimeout: 5 * time.Second,
	}

	log.Printf("Server is running on port %s", port)
	log.Fatal(srv.ListenAndServe())
}
