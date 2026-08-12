package main

import (
	"fmt"
)


type Payer interface {
	// Pay(amount int64) bool	// if returning anything
	Pay(amount int64) error
	Refund(amount int64, account string) error
}

type PaymentService struct{
	gateway Payer
}

// Constructor for payment service
func NewPaymentService(gateway Payer) *PaymentService{
	return &PaymentService{
		gateway: gateway,
	}
}

// open close principle
func (p *PaymentService) ProcessPayment(amount int64) error{
	// razorpayPaymentGw := razorpay{}
	// stripePaymentGw := stripe{}
	return p.gateway.Pay(amount)
}

func (p *PaymentService) ProcessRefund(amount int64, account string) error {
	return  p.gateway.Refund(amount, account)
}

// razorpay implementation
type razorpay struct{}

func (r razorpay) Pay(amount int64) error{
	// logic to make payment
	fmt.Println("making payment using razorpay", amount)
	return nil
}

func (r razorpay) Refund(amount int64, account string) error{
	fmt.Printf("Refund is done of amount, %d for %s\n", amount, account)
	return nil
}

// stripe implementation
type stripe struct{}

func (s stripe) Pay(amount int64) error{
	// logic to make payment
	fmt.Println("making payment using stripe", amount)
	return nil
}

func (s stripe) Refund(amount int64, account string) error{
	fmt.Printf("Refund is done of amount, %d for %s\n", amount, account)
	return nil
}

// fake/mock implementation for testing
type fakePayment struct{}

func (f fakePayment) Pay(amount int64) error{
	fmt.Println("Making fake payemnt for testing purpose.")
	return nil
}

func (f fakePayment) Refund(amount int64, account string) error{
	fmt.Printf("Refund is done of amount, %d for %s\n", amount, account)
	return nil
}

// PayPal implementation
type PayPal struct{}

func (p PayPal) Pay(amount int64) error {
	fmt.Println("Making payment using paypal", amount)
	return nil
}

func (p PayPal) Refund(amount int64, account string) error{
	fmt.Printf("Refund is done of amount, %d for %s\n", amount, account)
	return nil
}

func main() {
	// razorpayGw := razorpay{}
	// stripeGw := stripe{}
	paypalGw := PayPal{}
	// fakeGw := fakePayment{}

	// myPayment := paymentService{
	// 	// gateway: razorpayGw,
	// 	gateway: stripeGw,
	// 	// gateway: fakeGw,
	// }

	paymentService := NewPaymentService(paypalGw)

	if err := paymentService.ProcessPayment(10000); err != nil {
		fmt.Println("Payment failed", err)
	}

	if err := paymentService.ProcessRefund(10000, "Prashant"); err != nil {
		fmt.Println("Refund failed", err)
	}

	// myPayment.makePayment(100)
}