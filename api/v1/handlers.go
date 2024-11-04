package v1

import "net/http"

func login(repo *Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func register(repo *Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func getAllProducts(repo *Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func getProductByID(repo *Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func createProduct(repo *Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func updateProduct(repo *Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func updateOrderStatus(repo *Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func deleteProduct(repo *Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func getOrders(repo *Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func getOrderByID(repo *Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func cancelOrder(repo *Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func createOrder(repo *Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
