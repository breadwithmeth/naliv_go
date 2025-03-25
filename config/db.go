package config

import (
	"database/sql"
	"fmt"
	"log"

)

func GetDBConnection() (*sql.DB, error) {
	// Настройки подключения к MySQL
	user := "root"
	password := "password"
	host := "localhost"
	port := 3306
	database := "naliv_go"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, database)

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
