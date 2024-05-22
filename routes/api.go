package routes

import (
	handlers "myproject/Handlers"
	"myproject/middleware"
	"net/http"
)

func SetAPI() http.Handler {
	mux := http.NewServeMux()

	// Route for User
	mux.Handle("POST /api/v1/user", middleware.MiddlewareBasicAuth(http.HandlerFunc(handlers.StoreUser)))
	mux.Handle("GET /api/v1/user", middleware.MiddlewareBasicAuth(http.HandlerFunc(handlers.GetUsers)))
	mux.Handle("PUT /api/v1/user/{id}", middleware.MiddlewareBasicAuth(http.HandlerFunc(handlers.UpdateUser)))
	mux.Handle("DELETE /api/v1/user/{id}", middleware.MiddlewareBasicAuth(http.HandlerFunc(handlers.DeleteUser)))

	// Route for Product
	mux.Handle("POST /api/v1/product", middleware.MiddlewareBasicAuth(http.HandlerFunc(handlers.StoreProduct)))
	mux.Handle("GET /api/v1/product", middleware.MiddlewareBasicAuth(http.HandlerFunc(handlers.GetProduct)))
	mux.Handle("PUT /api/v1/product/{id}", middleware.MiddlewareBasicAuth(http.HandlerFunc(handlers.UpdateProduct)))
	mux.Handle("DELETE /api/v1/product/{id}", middleware.MiddlewareBasicAuth(http.HandlerFunc(handlers.DeleteProduct)))

	return mux
}
