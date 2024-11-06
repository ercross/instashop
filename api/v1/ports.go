package v1

import (
	"instashop/api/model"
)

// Repository provides a data storage client to manipulate data on a given database server.
//
// Concrete implementations of Repository should wrap errors generated
// as a result of wrong data input from user with api.ErrInvalidUserInput
// to enable handlers propagate the errors effectively.
// Errors not wrapped with api.ErrInvalidUserInput will be considered an internal error
type Repository interface {
	ValidateCredentials(email, password string) (model.User, error)
	Register(email, password string) (uint, error)
	FetchAllProducts() ([]model.Product, error)
	FetchProductByID(id uint) (model.Product, error)
	CreateProduct(product model.Product) (id uint, err error)
	UpdateProduct(product model.Product) error
	UpdateOrderStatus(status model.OrderStatus, orderID uint) error
	DeleteProduct(id uint) error
	FetchUserOrders(userID uint) ([]model.Order, error)
	FetchOrderByID(id uint) (model.Order, error)

	// CancelOrder updates the order status to "Cancelled" if the order is in a "Pending" state
	CancelOrder(id uint) error
	CreateOrder(order model.Order) (id uint, err error)
}
