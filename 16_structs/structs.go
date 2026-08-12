package main

import (
	"fmt"
	"time"
)

type order struct {
	ID          string
	AmountCents int64 // Storing money as integer cents avoids floating-point inaccuracies
	Status      string
	createdAt   time.Time // nano sec precision
}


// reciever type
func (o *order) changeStatus(status string) {
	o.Status = status
}

func (o order) getAmount () int64{
	return o.AmountCents
}

func NewOrder(id string, amountCents int64, status string) *order {
	return &order{
		ID:          id,
		AmountCents: amountCents,
		Status:      status,
		createdAt:   time.Now(),
	}
}

// like classes in other languages we use here struct
func main() {
	// creating struct instance
	// if you don't set a field then default value will be zero value
	// int => 0, float => 0, string => "", bool => false
	myOrder := order{
		ID:          "1",
		AmountCents: 5000,
		Status:      "received",
	}
	
	myOrder2 := order{
		ID:          "2",
		AmountCents: 3090,
		Status:      "pending",
		createdAt:   time.Now(),
	}

	order1 := NewOrder("3", 3445, "received")
	order2 := NewOrder("4", 3100, "received")

	order1.changeStatus("confirmed")

	myOrder.createdAt = time.Now()

	fmt.Println("Order struct - 1:", myOrder)
	fmt.Println("Order struct - 2:", myOrder2)

	fmt.Printf("Order 1: %+v\n", order1) // Order 1: {ID:3 AmountCents:3445 Status:confirmed createdAt:{wall:14021728866650331272 ext:1 loc:0x7ff6eeaf74a0}}

	fmt.Printf("Order 2: %+v\n", order2) // Order 2: {ID:4 AmountCents:3100 Status:received createdAt:{wall:14021728242842607628 ext:1 loc:0x7ff77a9874a0}}

	fmt.Println(order2.getAmount())
}
