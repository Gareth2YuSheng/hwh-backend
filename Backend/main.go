package main

import (
	"fmt"
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
	logInfo("Loading env variables")
	//Load .env file
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		logFatal("PORT NOT FOUND", nil)
	}
	databaseURL := os.Getenv("DB_URL")
	if databaseURL == "" {
		logFatal("DB_URL NOT FOUND", nil)
	}
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		logFatal("JWT_SECRET NOT FOUND", nil)
	}

	logInfo("Connecting to DB")
	//Connect to DB
	store, err := NewPGStore(databaseURL)
	if err != nil {
		logFatal("UNABLE TO CONNECT TO DATABASE", err)
	}
	apiCfg := APIConfig{
		DB:        store,
		JwtSecret: jwtSecret,
	}
	//Init DB
	if err := store.dbInit(); err != nil {
		logFatal("UNABLE TO INITIALIZE DATABASE", err)
	}

	logInfo("Creating Routers")
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
	tagRouter := chi.NewRouter()

	//Account Routes
	accountRouter.Get("/user", apiCfg.middlewareAuth(apiCfg.handlerGetUserData))
	accountRouter.Post("/register", apiCfg.handlerCreateUser)
	accountRouter.Post("/login", apiCfg.handlerLogin)

	//Thread Routes
	threadRouter.Get("/all", apiCfg.middlewareAuth(apiCfg.handlerGetAllThreads))
	threadRouter.Get("/{threadID}/details", apiCfg.middlewareAuth(apiCfg.handlerGetTheadDetails))
	threadRouter.Post("/create", apiCfg.middlewareAuth(apiCfg.handlerCreateThread))
	threadRouter.Put("/{threadID}/update", apiCfg.middlewareAuth(apiCfg.handlerUpdateThread))
	threadRouter.Delete("/{threadID}/delete", apiCfg.middlewareAuth(apiCfg.handlerDeleteThread))

	//Tag Routes
	tagRouter.Get("/all", apiCfg.middlewareAuth(apiCfg.handlerGetAllTags))
	tagRouter.Post("/create", apiCfg.middlewareAuth(apiCfg.handlerCreateTag))

	//Comment Routes
	commentRouter.Get("/{threadID}", apiCfg.middlewareAuth(apiCfg.handlerGetAllComments))
	commentRouter.Post("/{threadID}/create", apiCfg.middlewareAuth(apiCfg.handlerCreateComment))
	commentRouter.Put("/{commentID}/update", apiCfg.middlewareAuth(apiCfg.handlerUpdateComment))
	commentRouter.Put("/{commentID}/answer", apiCfg.middlewareAuth(apiCfg.handlerMarkCommentAsAnswer))
	commentRouter.Put("/{commentID}/vote", apiCfg.middlewareAuth(apiCfg.handlerVoteComment))
	commentRouter.Delete("/{commentID}/delete", apiCfg.middlewareAuth(apiCfg.handlerDeleteComment))
	commentRouter.Delete("/{commentID}/unvote", apiCfg.middlewareAuth(apiCfg.handlerUnVoteComment))

	router.Mount("/account", accountRouter)
	router.Mount("/thread", threadRouter)
	router.Mount("/comment", commentRouter)
	router.Mount("/tag", tagRouter)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	logInfo(fmt.Sprintf("Server starting on port: %v", port))
	err = server.ListenAndServe()
	if err != nil {
		logFatal("FAILED TO START SERVER", err)
	}
}
