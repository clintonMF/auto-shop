package shop

import "fmt"

type Accessory struct {
	id        string
	name      string
	brandName string
	quantity  int
	price     float64
}

func AddAccessory(name, brandName string, quantity int, price float64) Accessory {
	id := GenUUID()
	return Accessory{
		id:        id,
		name:      name,
		brandName: brandName,
		quantity:  quantity,
		price:     price,
	}
}

func (A *Accessory) Name() string {
	return fmt.Sprintf("%s-%s", A.brandName, A.name)
}

func (A *Accessory) Quantity() int {
	return A.quantity
}

func (A *Accessory) Price() float64 {
	return A.price
}

func (A *Accessory) Sell(num int) {
	A.quantity = A.quantity - num
}

func (A *Accessory) IsFinished() bool {
	return A.quantity > 1
}
