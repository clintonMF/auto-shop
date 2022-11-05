package main

import (
	"fmt"

	"github.com/clintonMF/auto-shop/shop"
)

func main() {
	fmt.Println("i am here")

	// shop.Hi()
	volvo := shop.AddCar("Toyota", "Camry", "black", 10, 500000)
	shop.Display(&volvo)
	shop.AddItem(volvo)

}
