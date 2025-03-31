package transport

import (
	"net/http"

	"github.com/breadwithmeth/naliv_go/internal/handlers"
	service "github.com/breadwithmeth/naliv_go/internal/services"
)

func NewRouter(userService *service.UserService, itemService *service.ItemService) *http.ServeMux {
	router := http.NewServeMux()

	// Создаём обработчики
	userHandler := handlers.NewUserHandler(userService)
	itemHandler := handlers.NewItemHandler(itemService)
	router.HandleFunc("GET /users", userHandler.GetAllUsersHandler)
	// router.HandleFunc("GET /user/{id}", userHandler.GetUserHandler)
	// router.HandleFunc("POST /user", userHandler.CreateUserHandler)
	// router.HandleFunc("PUT /user/{id}", userHandler.UpdateUserHandler)
	// router.HandleFunc("DELETE /user/{id}", userHandler.DeleteUserHandler)

	router.HandleFunc("POST /items", itemHandler.GetAllItemsHandler)
	// Обработчик 404
	// router.HandleFunc("/{...}", func(w http.ResponseWriter, r *http.Request) {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	w.Write([]byte("404 - Page Not Found"))
	// })

	return router
}
