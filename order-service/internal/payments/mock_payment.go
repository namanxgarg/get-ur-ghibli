package payments

import "errors"

func ProcessPayment(amount int) error {
    // In production, you'd integrate with a real payment gateway
    // We'll just succeed if amount > 0
    if amount <= 0 {
        return errors.New("invalid payment amount")
    }
    // Payment succeeds
    return nil
}
