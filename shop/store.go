package shop

import "fmt"

type Store struct {
	product map[string]Product
}

// My store
var myStore map[string]Product

func (s Store) NumberOfProduct() {
	num := 0
	for i := 0; i < len(myStore); i++ {
		num++
	}

	fmt.Printf("You have %v unsold products in this store", num)
}

func (s Store) AddItem(item Product) {
	id := GenLongUUID()

	myStore[id] = item
}
