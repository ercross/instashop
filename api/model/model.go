package model

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

// ErrInvalidUserInput is a wrapper for all errors generated as a result of invalid user input
var ErrInvalidUserInput = errors.New("user input error")

type OrderStatus int8

const (
	OrderStatusUnknown OrderStatus = iota
	OrderStatusPending
	OrderStatusConfirmed
	OrderStatusShipped
	OrderStatusDelivered
	OrderStatusCanceled
	OrderStatusReturned
	OrderStatusRefunded
	OrderStatusFailed
)

// User represents a user in the e-commerce system.
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Email     string    `json:"email" gorm:"unique;not null" sql:"type:varchar(100)"`
	Password  string    `json:"-" gorm:"not null" sql:"type:varchar(255)"`
	IsAdmin   bool      `json:"-" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// Product represents a product in the e-commerce system.
type Product struct {
	ID          uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string         `json:"name" gorm:"not null" sql:"type:varchar(100)"`
	Description string         `json:"description" sql:"type:text"`
	Price       float64        `json:"price" gorm:"not null"`
	Quantity    int            `json:"quantity" gorm:"not null"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"deleted_at"`
}

// Order represents an order placed by a user.
type Order struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint           `json:"user_id" gorm:"not null"`
	Status    OrderStatus    `json:"status" gorm:"not null" sql:"type:int;default:1"`
	Total     float64        `json:"total" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"deleted_at"`
}

// OrderItem represents the items within an order, linking products to orders.
type OrderItem struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderID   uint           `json:"order_id" gorm:"not null"`
	ProductID uint           `json:"product_id" gorm:"not null"`
	Quantity  int            `json:"quantity" gorm:"not null"`
	Price     float64        `json:"price" gorm:"not null"` // Price at the time of the order
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"deleted_at"`
}
