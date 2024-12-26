package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

type APIConfig struct {
	DB        *PGStore
	JwtSecret string
}

func main() {
	//Load .env file
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT NOT FOUND")
	}
	databaseURL := os.Getenv("DB_URL")
	if databaseURL == "" {
		log.Fatal("DB_URL NOT FOUND")
	}
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET NOT FOUND")
	}

	//Connect to DB
	// store, err := NewPGStore(databaseURL)
	// if err != nil {
	// 	log.Fatalf("UNABLE TO CONNECT TO DATABASE: %v", err)
	// }
	apiCfg := APIConfig{
		// DB:        store,
		JwtSecret: jwtSecret,
	}
	// //Init DB
	// if err := store.Init(); err != nil {
	// 	log.Fatal(err)
	// }

	//ROUTERS
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	accountRouter := chi.NewRouter()
	threadRouter := chi.NewRouter()
	commentRouter := chi.NewRouter()

	//router.Get("/healthz", handlerTest)
	accountRouter.Post("/register", apiCfg.handlerCreateAccount)

	router.Mount("/account", accountRouter)
	router.Mount("/thread", threadRouter)
	router.Mount("/comment", commentRouter)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	log.Printf("Server starting on port %v", port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
