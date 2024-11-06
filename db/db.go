package db

import (
	"errors"
	"fmt"
	"instashop/api/model"

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

	err = db.AutoMigrate(&model.User{}, &model.Product{}, &model.Order{}, &model.OrderItem{})
	if err != nil {
		return nil, err
	}
	if err = seedAdminAccount(db); err != nil {
		return nil, err
	}
	return &DB{client: db}, nil
}

// seedAdminAccount creates an admin account to database if it does not already exist.
//
// Note that this is not a recommended approach to seed the database with an admin account
func seedAdminAccount(db *gorm.DB) error {

	email := "admin@instashop.com"
	password := "admin123"

	// Check if an admin account already exists
	var admin model.User
	if err := db.Where("email = ? AND is_admin = ?", email, true).First(&admin).Error; err == nil {

		return nil // Admin already exists, no need to create
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash admin password: %w", err)
	}

	newAdmin := model.User{
		Email:    email,
		Password: string(bytes),
		IsAdmin:  true,
	}
	if err := db.Create(&newAdmin).Error; err != nil {
		return fmt.Errorf("failed to create admin account: %w", err)
	}

	return nil
}

func (db *DB) ValidateCredentials(email, password string) (model.User, error) {
	var user model.User
	err := db.client.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, model.ErrInvalidUserInput
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

	user := model.User{
		Email:    email,
		Password: string(hashedPassword),
	}

	if err = db.client.Create(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return -1, fmt.Errorf("user already exists: %w", model.ErrInvalidUserInput)
		}
		return -1, fmt.Errorf("failed to create new user: %w", err)
	}

	return user.ID, nil
}

func (db *DB) FetchAllProducts() ([]model.Product, error) {
	var products []model.Product

	if err := db.client.Find(&products).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, gorm.ErrEmptySlice) {
			return products, nil
		}
		return nil, fmt.Errorf("error fetch products: %w", err)
	}

	return products, nil
}

func (db *DB) FetchProductByID(id uint) (model.Product, error) {
	var product model.Product
	if err := db.client.First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Product{}, fmt.Errorf("product not found: %w", model.ErrInvalidUserInput)
		}
		return product, fmt.Errorf("error fetching products: %w", err)
	}
	return product, nil
}

func (db *DB) CreateProduct(product model.Product) (uint, error) {
	if err := db.client.Create(&product).Error; err != nil {
		return -1, fmt.Errorf("failed to create product: %w", err)
	}
	return product.ID, nil
}

func (db *DB) UpdateProduct(product model.Product) error {
	if err := db.client.Save(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("product not found: %w", model.ErrInvalidUserInput)
		}
		return fmt.Errorf("failed to update product: %w", err)
	}
	return nil
}

func (db *DB) UpdateOrderStatus(status model.OrderStatus, orderID uint) error {
	var order model.Order
	err := db.client.Model(&order).Where("id = ?", orderID).Update("status", status).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("order does not exist: %w", model.ErrInvalidUserInput)
		}
		return fmt.Errorf("failed to update order: %w", err)
	}
	return err
}

func (db *DB) DeleteProduct(id uint) error {
	if err := db.client.Delete(&model.Product{}, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("product not found: %w", model.ErrInvalidUserInput)
		}
		return fmt.Errorf("failed to delete product: %w", err)
	}

	return nil
}

func (db *DB) FetchUserOrders(userID uint) ([]model.Order, error) {
	var orders []model.Order
	if err := db.client.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, gorm.ErrEmptySlice) {
			return orders, nil
		}
		return nil, fmt.Errorf("error fetching orders %d: %w", userID, err)
	}
	return orders, nil
}

// FetchOrderByID retrieves a single order by its ID
func (db *DB) FetchOrderByID(id uint) (model.Order, error) {
	var order model.Order
	if err := db.client.First(&order, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Order{}, fmt.Errorf("order not found: %w", model.ErrInvalidUserInput)
		}
		return order, fmt.Errorf("error fetching order: %w", err)
	}
	return order, nil
}

func (db *DB) CancelOrder(id uint) error {
	var order model.Order
	if err := db.client.First(&order, id).Error; err != nil {
		return err
	}

	if order.Status != model.OrderStatusPending {
		return errors.New("only pending orders can be cancelled")
	}

	order.Status = model.OrderStatusCanceled
	return db.client.Save(&order).Error
}

func (db *DB) CreateOrder(order model.Order) (uint, error) {
	err := db.client.Create(&order).Error
	return order.ID, err
}
