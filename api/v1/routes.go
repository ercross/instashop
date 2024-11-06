package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func AddRoutes(mux *chi.Mux, repo Repository) {
	mux.Use(middleware.AllowContentType("application/json"))
	mux.Use(authMiddleware)

	mux.Mount("/", adminRoutes(repo))
	mux.Mount("/auth", authenticationRoutes(repo))
	mux.Mount("/order", orderRoutes(repo))
}

func authenticationRoutes(repo Repository) http.Handler {
	r := chi.NewRouter()
	r.Post("/login", login(repo))
	r.Post("/register", register(repo))
	return r
}

func adminRoutes(repo Repository) http.Handler {
	r := chi.NewRouter()

	r.Use(adminOnlyMiddleware)

	r.Get("/products", getAllProducts(repo))
	r.Get("/product/{id}", getProductByID(repo))

	r.Post("/product", createProduct(repo))

	r.Put("/product/", updateProduct(repo))
	r.Put("/orders", updateOrderStatus(repo))
	r.Delete("/products", deleteProduct(repo))

	return r
}

func orderRoutes(repo Repository) http.Handler {
	r := chi.NewRouter()

	r.Get("/", getOrders(repo))
	r.Get("/{id}", getOrderByID(repo))

	r.Put("/cancel", cancelOrder(repo))

	r.Post("/new", createOrder(repo))

	return r
}
