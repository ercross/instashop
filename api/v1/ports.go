package v1

import "instashop/api"

// Repository provides a data storage client to manipulate data on a given database server.
//
// Concrete implementations of Repository should wrap errors generated
// as a result of wrong data input from user with api.ErrInvalidUserInput
// to enable handlers propagate the errors effectively.
// Errors not wrapped with api.ErrInvalidUserInput will be considered an internal error
type Repository interface {
	Login(email, password string) error
	Register(email, password string) (uint, error)
	FetchAllProducts() ([]api.Product, error)
	FetchProductByID(id uint) (api.Product, error)
	CreateProduct(product api.Product) (id uint, err error)
	UpdateProduct(product api.Product) error
	UpdateOrderStatus(status api.OrderStatus, orderID uint) error
	DeleteProduct(id uint) error
	FetchUserOrders(userID uint) ([]api.Order, error)
	FetchOrderByID(id uint) (api.Order, error)

	// CancelOrder updates the order status to "Cancelled" if the order is in a "Pending" state
	CancelOrder(id uint) error
	CreateOrder(order api.Order) (id uint, err error)
}
