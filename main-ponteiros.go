package main

import "fmt"

type Order struct {
	ID       string
	Price    float64
	Quantity int
}

func NewOrder() *Order {
	return &Order{}
}

func (o *Order) setPrice(price float64) {
	o.Price = price
	fmt.Println("Price dentro do setPrice:", o.Price)
}

func (o Order) getTotal() float64 {
	return o.Price * float64(o.Quantity)
}

func main() {
	order := NewOrder()
	order.ID = "123"
	order.Price = 10.0
	order.Quantity = 5

	order.setPrice(20.0)
	fmt.Println("Preço total: ", order.getTotal())

	order2 := order
	order2.Price = 30.0
	fmt.Println("Preço total atualizado order2: ", order2.getTotal())
	fmt.Println("Preço total atualizado order: ", order.getTotal())
}
