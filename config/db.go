package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetDBConnection() (*sql.DB, error) {
	// Загружаем переменные окружения из .env файла
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %s\n", err.Error())
		return nil, err
	}

	// Настройки подключения к MySQL из переменных окружения
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("Error opening database: %s\n", err.Error())
		return nil, err
	}

	// Проверяем соединение
	if err := db.Ping(); err != nil {
		log.Printf("Error connecting to database: %s\n", err.Error())
		return nil, err
	}

	log.Println("Successfully connected to MySQL database")
	return db, nil
}
