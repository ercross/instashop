package v1

import "net/http"

func login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func getAllProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func getProductByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func createProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func updateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func updateOrderStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func deleteProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func getOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func getOrderByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func cancelOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func createOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
