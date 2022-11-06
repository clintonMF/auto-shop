package shop

import "fmt"

type car struct {
	id        string
	brandName string
	model     string
	color     string
	quantity  int
	price     float64
}

func AddCar(brandName, model, color string, quantity int, price float64) car {
	id := GenUUID()
	return car{
		id:        id,
		brandName: brandName,
		model:     model,
		color:     color,
		quantity:  quantity,
		price:     price,
	}
}

func (c *car) Name() string {
	return fmt.Sprintf("%s-%s [%s]", c.brandName, c.model, c.color)
}

func (c *car) Quantity() int {
	return c.quantity
}

func (c *car) Price() float64 {
	return c.price
}

func (c *car) Sell(num int) {
	c.quantity = c.quantity - num
}

func (c *car) IsFinished() bool {
	return c.quantity > 1
}
