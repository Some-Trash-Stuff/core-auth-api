package main

import (
	"log"
	"net/http"

	"core-auth-api/configs"
	v1 "core-auth-api/internal/v1"
)

func main() {
	cfg := configs.Load()

	router := v1.NewRouter()

	log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}
