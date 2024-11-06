package v1

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"instashop/api/model"
	"log"
	"net/http"
	"strconv"
)

var errInternalServerError = errors.New("Internal Server Error")

func login(repo Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		user, err := repo.ValidateCredentials(req.Email, req.Password)
		if err != nil {
			if errors.Is(err, model.ErrInvalidUserInput) {
				http.Error(w, "Invalid credentials", http.StatusUnauthorized)
				return
			}

			log.Printf("Error validating login credentials: %v", err)
			http.Error(w, "Login failed. Please try again", http.StatusInternalServerError)
			return
		}

		token, err := generateToken(user.ID, user.IsAdmin)
		if err != nil {
			log.Printf("Error generating token: %v", err)
			http.Error(w, "Login failed. Please try again", http.StatusInternalServerError)
			return
		}

		sendJSONResponse(w, http.StatusOK, map[string]string{"token": token})
	}
}

func register(repo Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		userID, err := repo.Register(req.Email, req.Password)
		if err != nil {
			if errors.Is(err, model.ErrInvalidUserInput) {
				http.Error(w, "Email already in use", http.StatusConflict)
				return
			}
			log.Printf("Error registering user: %v", err)
			http.Error(w, "Failed to register user", http.StatusInternalServerError)
			return
		}

		sendJSONResponse(w, http.StatusCreated, map[string]uint{"user_id": userID})
	}
}

func getAllProducts(repo Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := repo.FetchAllProducts()
		if err != nil {
			http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
			return
		}

		sendJSONResponse(w, http.StatusOK, products)
	}
}

func getProductByID(repo Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}

		product, err := repo.FetchProductByID(uint(id))
		if err != nil {
			if errors.Is(err, model.ErrInvalidUserInput) {
				http.Error(w, "Product not found", http.StatusNotFound)
				return
			}
			log.Printf("Error fetching product: %v", err)
			http.Error(w, errInternalServerError.Error(), http.StatusInternalServerError)
			return
		}

		sendJSONResponse(w, http.StatusOK, product)
	}
}

func createProduct(repo Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var product model.Product
		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		id, err := repo.CreateProduct(product)
		if err != nil {
			http.Error(w, "Failed to create product", http.StatusInternalServerError)
			return
		}

		sendJSONResponse(w, http.StatusCreated, map[string]interface{}{"id": id})
	}
}

func updateProduct(repo Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var product model.Product
		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		err := repo.UpdateProduct(product)
		if err != nil {
			if errors.Is(err, model.ErrInvalidUserInput) {
				http.Error(w, "Invalid product id", http.StatusBadRequest)
				return
			}
			log.Printf("Error updating product: %v", err)
			http.Error(w, errInternalServerError.Error(), http.StatusInternalServerError)
			return
		}

		sendJSONResponse(w, http.StatusOK, nil)
	}
}

func updateOrderStatus(repo Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Status model.OrderStatus `json:"status"`
			ID     uint              `json:"id"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		err := repo.UpdateOrderStatus(req.Status, req.ID)
		if err != nil {
			if errors.Is(err, model.ErrInvalidUserInput) {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			log.Printf("Error updating order status: %v", err)
			http.Error(w, errInternalServerError.Error(), http.StatusInternalServerError)
			return
		}

		sendJSONResponse(w, http.StatusOK, nil)
	}
}

func deleteProduct(repo Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}

		err = repo.DeleteProduct(uint(id))
		if err != nil {
			if errors.Is(err, model.ErrInvalidUserInput) {
				http.Error(w, "Product not found", http.StatusNotFound)
				return
			}
			http.Error(w, errInternalServerError.Error(), http.StatusInternalServerError)
			return
		}

		sendJSONResponse(w, http.StatusOK, nil)
	}
}

func getOrders(repo Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("user_id").(uint)

		orders, err := repo.FetchUserOrders(userID)
		if err != nil {
			log.Printf("Error fetching orders for user %d: %v", userID, err)
			http.Error(w, "Failed to fetch orders", http.StatusInternalServerError)
			return
		}

		sendJSONResponse(w, http.StatusOK, orders)
	}
}

func getOrderByID(repo Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid order ID", http.StatusBadRequest)
			return
		}

		order, err := repo.FetchOrderByID(uint(id))
		if err != nil {
			if errors.Is(err, model.ErrInvalidUserInput) {
				http.Error(w, "Order not found", http.StatusNotFound)
				return
			}

			log.Printf("Error fetching order: %v", err)
			http.Error(w, errInternalServerError.Error(), http.StatusInternalServerError)
			return
		}

		sendJSONResponse(w, http.StatusOK, order)
	}
}

func cancelOrder(repo Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid order ID", http.StatusBadRequest)
			return
		}

		err = repo.CancelOrder(uint(id))
		if err != nil {
			if errors.Is(err, model.ErrInvalidUserInput) {
				http.Error(w, "Order not found", http.StatusBadRequest)
				return
			}
			log.Printf("Error fetching order: %v", err)
			http.Error(w, errInternalServerError.Error(), http.StatusInternalServerError)
			return
		}

		sendJSONResponse(w, http.StatusOK, nil)
	}
}

func createOrder(repo Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("user_id").(uint)

		var order model.Order
		if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		order.UserID = userID
		id, err := repo.CreateOrder(order)
		if err != nil {
			http.Error(w, "Failed to create order", http.StatusInternalServerError)
			return
		}

		sendJSONResponse(w, http.StatusCreated, map[string]uint{"order_id": id})
	}
}

func sendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)

	if err := encoder.Encode(data); err != nil {
		http.Error(w, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
	}
}
