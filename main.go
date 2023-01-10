package main

import (
	"fmt"

	"github.com/clintonMF/auto-shop/shop"
)

func main() {
	fmt.Println("WELCOME TO JOHN'S AUTO SHOP")
	shop.Newline(1)
	// creating store
	store := shop.NewStore()

	// Creating cars
	fmt.Println("Creating cars .......")
	modelS := shop.AddCar("Tesla", "ModelS", "Red", 50, 2500000)
	modelY := shop.AddCar("Tesla", "ModelY", "Green", 5, 6000000)
	modelSb := shop.AddCar("Tesla", "ModelS", "black", 30, 2700000)
	camry := shop.AddCar("Toyata", "Camary", "black", 10, 500000)

	// Adding items to store
	fmt.Println("Adding items to store ..........")
	store.AddItem(&camry)
	store.AddItem(&modelS)
	store.AddItem(&modelSb)
	store.AddItem(&modelY)
	// Getting the various number of products added to store
	store.NumberOfProduct()

	// list all item in store
	fmt.Println("List of Items in store")
	store.DisplayAllItem()

	// Selling items
	fmt.Println("selling items ......")
	shop.Newline(1)
	store.Sell(2, 6)
	store.Sell(1, 4)
	store.Sell(3, 15)
	store.Sell(4, 2)

	// List of items after selling
	fmt.Println("List of items after selling")
	store.DisplayAllItem()

	// Log book/Sales history
	store.ItemsLog()
}
