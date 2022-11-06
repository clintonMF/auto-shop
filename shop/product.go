package shop

import "fmt"

type Product interface {
	Name() string
	Quantity() int
	Price() float64
	Sell(qty int)
	IsFinished() bool
}

func Display(p Product) {
	name := p.Name()
	quantity := p.Quantity()
	price := p.Price()

	fmt.Printf("Name - %s \nQuantity - %v \nPrice - %.2f", name, quantity,
		price)
	Newline(2)
}
