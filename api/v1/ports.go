package v1

import "instashop/api"

type Repository interface {
	Login(email, password string) error
	Register(email, password string) error
	FetchAllProducts() ([]api.Product, error)
	FetchProductByID(id uint) (api.Product, error)
	CreateProduct(product api.Product) (id uint, err error)
	UpdateProduct(product api.Product) error
	UpdateOrderStatus(status api.OrderStatus) error
	DeleteProduct(id uint) error
	FetchUserOrders(userID uint) ([]api.Order, error)
	FetchOrderByID(id uint) (api.Order, error)
	CancelOrder(id uint) error
	CreateOrder(order api.Order) (id uint, err error)
}
