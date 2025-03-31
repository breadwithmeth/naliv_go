package app

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	// "github.com/breadwithmeth/naliv_go/internal/config"
	config "github.com/breadwithmeth/naliv_go/internal/config"
	"github.com/breadwithmeth/naliv_go/internal/database"
	"github.com/breadwithmeth/naliv_go/internal/repository"
	service "github.com/breadwithmeth/naliv_go/internal/services"
	transport "github.com/breadwithmeth/naliv_go/internal/transport/rest"
)

// Run запускает приложение
func Run() {
	ctx := context.Background()

	// Загружаем конфигурацию
	cfg := config.LoadConfig()
	if cfg == nil {
		log.Fatal("Failed to load configuration")
	}

	// Подключаемся к базе данных
	db, err := database.Connect(cfg.GetDSN())
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}
	defer db.Close()

	router := setupRouter(db)

	// Формируем адрес сервера
	serverAddr := "0.0.0.0:6150"

	// Запускаем сервер
	if err := startServer(ctx, serverAddr, router); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

// Инициализация маршрутов и зависимостей
func setupRouter(db *sql.DB) *http.ServeMux {
	// 1. Создаём репозитории
	userRepo := repository.NewUserRepository(db)
	itemRepo := repository.NewItemRepository(db)
	// 2. Создаём сервисы
	userService := service.NewUserService(userRepo)
	itemService := service.NewItemService(itemRepo)
	// 3. Создаём маршрутизатор с обработчиками
	router := transport.NewRouter(userService, itemService)

	return router
}

// Запуск HTTP-сервера
func startServer(ctx context.Context, address string, router *http.ServeMux) error {
	server := &http.Server{
		Addr:         address,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Канал для сигналов завершения
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Запуск сервера в горутине
	go func() {
		log.Printf("Starting server on %s...", address)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Ожидание сигнала завершения
	select {
	case <-done:
		log.Print("Server stopping...")
	case <-ctx.Done():
		log.Print("Server stopping...")
	}

	// Graceful shutdown
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		return err
	}

	log.Print("Server stopped")
	return nil
}
