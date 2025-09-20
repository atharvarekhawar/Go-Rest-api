package main

import (
	"fmt"
	"log"
	"net/http"
	"rest-api/internal/config"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
	cfg := config.SerializeConfig()

	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to server 3000"))
	})

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	fmt.Println("Server started")

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal("failed to start server")
	}

}
