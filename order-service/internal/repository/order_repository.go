package repository

import (
    "time"

    "github.com/example/get-ur-ghibli/order-service/internal/models"
    "github.com/jinzhu/gorm"
)

type OrderRepository interface {
    CreateOrder(o *models.Order) error
    GetOrder(id uint) (*models.Order, error)
    UpdateOrder(o *models.Order) error
}

type orderRepo struct {
    db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
    return &orderRepo{db: db}
}

func (r *orderRepo) CreateOrder(o *models.Order) error {
    return r.db.Create(o).Error
}

func (r *orderRepo) GetOrder(id uint) (*models.Order, error) {
    var order models.Order
    err := r.db.First(&order, id).Error
    if err != nil {
        return nil, err
    }
    return &order, nil
}

func (r *orderRepo) UpdateOrder(o *models.Order) error {
    return r.db.Save(o).Error
}

// Helper function (not strictly needed) to set order as Paid
func (r *orderRepo) MarkPaid(o *models.Order) error {
    now := time.Now()
    o.Status = models.StatusPaid
    o.PaidAt = &now
    return r.db.Save(o).Error
}
