package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "github.com/jinzhu/gorm"

    "github.com/example/get-ur-ghibli/order-service/internal/models"
    "github.com/example/get-ur-ghibli/order-service/internal/payments"
    "github.com/example/get-ur-ghibli/order-service/internal/repository"
)

func CreateOrderHandler(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req struct {
            Email     string `json:"email"`
            OrderType string `json:"orderType"`
            Address   string `json:"address"`
            ImageRef  string `json:"imageRef"`
        }
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            http.Error(w, "Bad request", http.StatusBadRequest)
            return
        }

        var amount int
        switch req.OrderType {
        case string(models.OrderTypeTenImages):
            amount = 100
        case string(models.OrderType3DModel):
            amount = 2000
        default:
            http.Error(w, "Invalid order type", http.StatusBadRequest)
            return
        }

        orderRepo := repository.NewOrderRepository(db)
        order := &models.Order{
            Email:     req.Email,
            OrderType: models.OrderType(req.OrderType),
            Amount:    amount,
            Address:   req.Address,
            ImageRef:  req.ImageRef,
        }

        if err := orderRepo.CreateOrder(order); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(order)
    }
}

func PayOrderHandler(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        orderIDStr := vars["orderID"]
        orderID, err := strconv.Atoi(orderIDStr)
        if err != nil {
            http.Error(w, "Invalid order ID", http.StatusBadRequest)
            return
        }

        orderRepo := repository.NewOrderRepository(db)
        order, err := orderRepo.GetOrder(uint(orderID))
        if err != nil {
            http.Error(w, "Order not found", http.StatusNotFound)
            return
        }

        if order.Status != models.StatusCreated {
            http.Error(w, "Order is not payable in current state", http.StatusConflict)
            return
        }

        // Mock Payment
        err = payments.ProcessPayment(order.Amount)
        if err != nil {
            http.Error(w, "Payment failed", http.StatusPaymentRequired)
            return
        }

        // Mark paid
        order.Status = models.StatusPaid
        if err := orderRepo.UpdateOrder(order); err != nil {
            http.Error(w, "Could not update order status", http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(order)
    }
}

func GetOrderHandler(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        orderIDStr := vars["orderID"]
        orderID, err := strconv.Atoi(orderIDStr)
        if err != nil {
            http.Error(w, "Invalid order ID", http.StatusBadRequest)
            return
        }

        orderRepo := repository.NewOrderRepository(db)
        order, err := orderRepo.GetOrder(uint(orderID))
        if err != nil {
            http.Error(w, "Order not found", http.StatusNotFound)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(order)
    }
}
