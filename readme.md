### ALT SCHOOL GOLANG EXAM

## question

John has just opened up his car selling shop, to sell different cars. He gets the cars he needs to sell from different people and they all bring it to him. 
He needs to manage the list of cars he has, attach a price to them, and put them on display to be sold, basically John needs an inventory to manage cars & to manage sales. For instance, 

-  needs to see the number of cars that are left to be sold 
- He needs to see the sum of the prices of the cars left
- He needs to see the number of cars he has sold
- Sum total of the prices of cars he has sold
- A list of orders that for the sales he made

Using the knowledge of OOP in Go, Build simple classes for the following “objects”

Car
Product
Store

The Car class can have any car attributes you can think of.

The Product class should have attributes of a product i.e (the product, quantity of the product in stock, price of the product). A car is a product of the store, but there can be other products so the attribute of the car can be promoted to the Product. The Product class should have methods to display a product, and a method to display the status of a product if it is still in stock or not.

The Store class should have attributes like

- Number of products in the store that are still up for sale
- Adding an Item to the store
- Listing all product items in the store
- Sell an item
- Show a list of sold items and the total price

## Project structure

```
project
│   go.mod
|   go.sum
|   main.go
│   README.md
│
└─── shop
        accessory.go
        car.go
        product.go
        store.go
        util.go
```