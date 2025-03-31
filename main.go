package main

import (
	"log"

	"github.com/breadwithmeth/naliv_go/internal/app"
)

func main() {
	// Инициализация репозитория
	app.Run()
	log.Println("Starting naliv_go application...")
}
