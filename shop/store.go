package shop

import "fmt"

type Order struct {
	itemId   int
	quantity int
}

type Store struct {
	myStore map[int]Product
	Orders  []Order
}

var logBook []Order // stores store data

// This function creates a new store
func NewStore() *Store {
	return &Store{
		myStore: make(map[int]Product),
		Orders:  []Order{},
	}
}

// This function tells the number of unique product in store
func (s Store) NumberOfProduct() {
	num := 0
	for i := 0; i < len(s.myStore); i++ {
		num++
	}

	fmt.Printf("You have %v types of products in this store", num)
	Newline(2)
}

// This function is used to add Unique item to the store
func (s Store) AddItem(item Product) {
	num := len(s.myStore)
	num++
	s.myStore[num] = item
}

/* This function is used to display all the items in the store
it also shows
- Total number of cars left
- Total sum total of cars left
*/
func (s Store) DisplayAllItem() {
	fmt.Println("Below is a list of all the items in store")
	Newline(1)
	fmt.Println("= ID ====== Name ======== Price === Quantity")
	var num_of_cars_left int
	var total_sum_of_cars_left float64
	for i, v := range s.myStore {
		total_sum_of_cars_left += float64(v.Quantity()) * v.Price()
		num_of_cars_left += v.Quantity()
		fmt.Printf("%v - %s - %.2f - %v", i, v.Name(), v.Price(), v.Quantity())
		Newline(1)
	}

	Newline(1)
	fmt.Printf("Total number of cars left  =>  %d\n", num_of_cars_left)
	fmt.Printf("Total sum total of cars left  =>  %.2f", total_sum_of_cars_left)
	Newline(2)
}

// this function is used to sell items in the store
func (s Store) Sell(itemId, qty int) {
	item := s.myStore[itemId]
	name := item.Name()
	if !item.IsFinished() {
		fmt.Printf("No item %s in stock", name)
	} else if item.Quantity() < qty {
		fmt.Println("Quantity left in stock is less than you requested \nquantity left ", item.Quantity())
	} else {
		item.Sell(qty)
		newOrder := Order{
			itemId:   itemId,
			quantity: qty,
		}
		s.Orders = append(s.Orders, newOrder)
	}

	logBook = append(logBook, s.Orders...)
}

/* This function is used to display a list of all items sold
it also shows
- Total number of cars sold
- Total sum of cars sold
*/
func (s Store) ItemsLog() {
	Newline(1)
	fmt.Println("SALES HISTORY")
	fmt.Println("Num ====== Name ===== Price ==== Qty = Total price")
	var total_sales_made float64
	var total_number_of_cars_sold int
	for i, v := range logBook {
		num := i + 1
		item := s.myStore[v.itemId]
		price := item.Price()
		qty := v.quantity
		totalPrice := float64(qty) * price

		fmt.Printf("%d - %v - %.2f, %v, %.2f", num, item.Name(), price, qty, totalPrice)
		total_sales_made += totalPrice
		total_number_of_cars_sold += qty
		Newline(1)
	}

	Newline(1)
	fmt.Printf("Total number of cars sold  =>  %d\n", total_number_of_cars_sold)
	fmt.Printf("Total sum of cars sold  =>  %.2f", total_sales_made)
}
