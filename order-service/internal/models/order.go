package models

import (
    "github.com/jinzhu/gorm"
    "time"
)

type OrderType string

const (
    OrderTypeTenImages  OrderType = "TEN_IMAGES"
    OrderType3DModel    OrderType = "3D_MODEL"
)

type OrderStatus string

const (
    StatusCreated   OrderStatus = "CREATED"
    StatusPaid      OrderStatus = "PAID"
    StatusShipping  OrderStatus = "SHIPPING"
    StatusDelivered OrderStatus = "DELIVERED"
)

type Order struct {
    gorm.Model
    Email      string      `gorm:"not null"` // which user made the order
    OrderType  OrderType   `gorm:"type:varchar(50);not null"`
    Amount     int         `gorm:"not null"`
    Status     OrderStatus `gorm:"type:varchar(50);not null;default:'CREATED'"`
    PaidAt     *time.Time
    Address    string
    ImageRef   string // if user chooses an image for 3D
}
