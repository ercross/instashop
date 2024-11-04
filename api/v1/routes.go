package v1

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func addRoutes(repo *Repository) {

}

func authenticationRoutes() http.Handler {
	r := chi.NewRouter()
	r.Post("/login", login())
	r.Post("/register", register())
	return r
}

func adminRoutes() http.Handler {
	r := chi.NewRouter()

	r.Get("/products", getAllProducts())
	r.Get("/product/{id}", getProductByID())

	r.Post("/product", createProduct())

	r.Put("/product/", updateProduct())
	r.Put("/orders", updateOrderStatus())
	r.Delete("/products", deleteProduct())

	return r
}

func orderRoutes() http.Handler {
	r := chi.NewRouter()

	r.Get("/orders", getOrders())
	r.Get("/order/{id}", getOrderByID())

	r.Put("/order/cancel", cancelOrder())

	r.Post("/orders", createOrder())

	return r
}
