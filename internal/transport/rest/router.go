package transport

import (
	"net/http"

	"github.com/breadwithmeth/naliv_go/internal/handlers"
	service "github.com/breadwithmeth/naliv_go/internal/services"
	"github.com/breadwithmeth/naliv_go/internal/transport/rest/middleware"
)

func NewRouter(userService *service.UserService, itemService *service.ItemService, categoryService *service.CategoryService) *http.ServeMux {
	router := http.NewServeMux()

	v1 := http.NewServeMux()
	// Применяем все middleware
	v1WithAuth := middleware.AuthMiddleware(userService)(v1)
	v1WithCORS := middleware.CORSMiddleware(v1WithAuth)
	v1WithMiddleware := middleware.LoggingMiddleware(v1WithCORS)
	router.Handle("/v1/", http.StripPrefix("/v1", v1WithMiddleware))

	// Создаём обработчики
	userHandler := handlers.NewUserHandler(userService)
	itemHandler := handlers.NewItemHandler(itemService)
	categoryHandler := handlers.NewCategoryHandler(categoryService)
	v1.HandleFunc("GET /users", userHandler.GetAllUsersHandler)
	// router.HandleFunc("GET /user/{id}", userHandler.GetUserHandler)
	// router.HandleFunc("POST /user", userHandler.CreateUserHandler)
	// router.HandleFunc("PUT /user/{id}", userHandler.UpdateUserHandler)
	// router.HandleFunc("DELETE /user/{id}", userHandler.DeleteUserHandler)

	v1.HandleFunc("POST /items", itemHandler.GetAllItemsHandler)

	v1.HandleFunc("POST /categories", categoryHandler.GetAllCategoriesHandler)
	// Обработчик 404
	// router.HandleFunc("/{...}", func(w http.ResponseWriter, r *http.Request) {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	w.Write([]byte("404 - Page Not Found"))
	// })

	return router
}
