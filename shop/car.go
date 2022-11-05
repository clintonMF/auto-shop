package shop

import "fmt"

type Car struct {
	brandName string
	model     string
	color     string
	quantity  int
	price     float64
}

func (c *Car) AddCar(brandName, model, color string, quantity int, price float64) Car {
	return Car{
		brandName: brandName,
		model:     model,
		color:     color,
		quantity:  quantity,
		price:     price,
	}
}

func (c *Car) Name() string {
	return fmt.Sprintf("%s-%s (%s)", c.brandName, c.model, c.color)
}

func (c *Car) Quantity() int {
	return c.quantity
}

func (c *Car) Price() float64 {
	return c.price
}
