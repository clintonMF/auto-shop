package shop

import "fmt"

type Product interface {
	Name() string
	Quantity() int
	Price() float64
}

func Display(p Product) {
	name := p.Name()
	quantity := p.Quantity()
	price := p.Price()

	fmt.Printf("Name - %s \nQuantity - %v \n Price - %.2f", name, quantity,
		price)
}
