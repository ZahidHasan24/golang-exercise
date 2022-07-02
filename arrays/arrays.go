//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	price int
	name  string
}

//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
func printResult(itemList [4]Product) {

	var totalItems, totalCost int
	for i := 0; i < len(itemList); i++ {
		item := itemList[i]
		totalCost += item.price
		if item.name != "" {
			totalItems += 1
		}
	}

	fmt.Println("Last item on the list: ", itemList[totalItems-1])
	fmt.Println("Total items: ", totalItems)
	fmt.Println("Total cost: ", totalCost)
}

func main() {

	//* Using an array, create a shopping list with enough room
	//  for 4 products
	shoppingList := [4]Product{
		{50, "shirt"},
		{70, "pant"},
		{40, "shoe"},
	}

	printResult(shoppingList)

	//* Add a fourth product to the list and print out the
	//  information again
	shoppingList[3] = Product{10, "sunglass"}

	printResult(shoppingList)

}
