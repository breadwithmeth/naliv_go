package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // Импортируем драйвер MySQL
)

func Connect(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn) // Измените на "postgres" для PostgreSQL
	if err != nil {
		return nil, err
	}

	// Проверяем соединение
	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Database connected successfully")
	return db, nil
}
