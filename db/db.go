package db

import (
	"errors"
	"fmt"
	"instashop/api"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	client *gorm.DB
}

func NewDB(dsn string) (*DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&api.User{}, &api.Product{}, &api.Order{}, &api.OrderItem{})
	if err != nil {
		return nil, err
	}

	return &DB{client: db}, nil
}

func (db *DB) ValidateCredentials(email, password string) (api.User, error) {
	var user api.User
	err := db.client.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, api.ErrInvalidUserInput
		}
		return user, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return user, errors.New("invalid credentials")
	}

	return user, nil
}

func (db *DB) Register(email, password string) (uint, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return -1, fmt.Errorf("failed to generate hash from password: %w", err)
	}

	user := api.User{
		Email:    email,
		Password: string(hashedPassword),
	}

	if err = db.client.Create(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return -1, fmt.Errorf("user already exists: %w", api.ErrInvalidUserInput)
		}
		return -1, fmt.Errorf("failed to create new user: %w", err)
	}

	return user.ID, nil
}

func (db *DB) FetchAllProducts() ([]api.Product, error) {
	var products []api.Product

	if err := db.client.Find(&products).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, gorm.ErrEmptySlice) {
			return products, nil
		}
		return nil, fmt.Errorf("error fetch products: %w", err)
	}

	return products, nil
}

func (db *DB) FetchProductByID(id uint) (api.Product, error) {
	var product api.Product
	if err := db.client.First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return api.Product{}, fmt.Errorf("product not found: %w", api.ErrInvalidUserInput)
		}
		return product, fmt.Errorf("error fetching products: %w", err)
	}
	return product, nil
}

func (db *DB) CreateProduct(product api.Product) (uint, error) {
	if err := db.client.Create(&product).Error; err != nil {
		return -1, fmt.Errorf("failed to create product: %w", err)
	}
	return product.ID, nil
}

func (db *DB) UpdateProduct(product api.Product) error {
	if err := db.client.Save(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("product not found: %w", api.ErrInvalidUserInput)
		}
		return fmt.Errorf("failed to update product: %w", err)
	}
	return nil
}

func (db *DB) UpdateOrderStatus(status api.OrderStatus, orderID uint) error {
	var order api.Order
	err := db.client.Model(&order).Where("id = ?", orderID).Update("status", status).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("order does not exist: %w", api.ErrInvalidUserInput)
		}
		return fmt.Errorf("failed to update order: %w", err)
	}
	return err
}

func (db *DB) DeleteProduct(id uint) error {
	if err := db.client.Delete(&api.Product{}, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("product not found: %w", api.ErrInvalidUserInput)
		}
		return fmt.Errorf("failed to delete product: %w", err)
	}

	return nil
}

func (db *DB) FetchUserOrders(userID uint) ([]api.Order, error) {
	var orders []api.Order
	if err := db.client.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, gorm.ErrEmptySlice) {
			return orders, nil
		}
		return nil, fmt.Errorf("error fetching orders %d: %w", userID, err)
	}
	return orders, nil
}

// FetchOrderByID retrieves a single order by its ID
func (db *DB) FetchOrderByID(id uint) (api.Order, error) {
	var order api.Order
	if err := db.client.First(&order, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return api.Order{}, fmt.Errorf("order not found: %w", api.ErrInvalidUserInput)
		}
		return order, fmt.Errorf("error fetching order: %w", err)
	}
	return order, nil
}

func (db *DB) CancelOrder(id uint) error {
	var order api.Order
	if err := db.client.First(&order, id).Error; err != nil {
		return err
	}

	if order.Status != api.OrderStatusPending {
		return errors.New("only pending orders can be cancelled")
	}

	order.Status = api.OrderStatusCanceled
	return db.client.Save(&order).Error
}

func (db *DB) CreateOrder(order api.Order) (uint, error) {
	err := db.client.Create(&order).Error
	return order.ID, err
}
