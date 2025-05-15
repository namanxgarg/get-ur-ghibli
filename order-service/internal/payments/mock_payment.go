package payments

import (
	"errors"
	"os"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
)

func ProcessPayment(amount int) error {
	// In production, you'd integrate with a real payment gateway
	// We'll just succeed if amount > 0
	if amount <= 0 {
		return errors.New("invalid payment amount")
	}
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	params := &stripe.PaymentIntentParams{
		Amount:        stripe.Int64(int64(amount * 100)), // amount in cents
		Currency:      stripe.String(string("usd")),
		PaymentMethod: stripe.String("pm_card_visa"), // test card
		Confirm:       stripe.Bool(true),
	}
	_, err := paymentintent.New(params)
	if err != nil {
		return err
	}
	return nil
}
