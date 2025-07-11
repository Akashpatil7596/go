package main

import "fmt"

// func main() {
// 	prices := []string{"A", "B"}
// 	fmt.Println(prices[0:])
// 	prices = append(prices, "C")
// 	fmt.Println(prices)
// }

// func main() {
// 	var productNames [2]string
// 	productNames = [2]string{"Fan"}
// 	prices := [4]float64{10.99, 19.99, 29.99, 39.99}
// 	productNames[0] = "as"
// 	fmt.Println(prices[len(prices)-1])
// 	fmt.Println(len(productNames))

// 	featuredPrice := prices[1:]
// 	highlightedPrice := featuredPrice[:1]
// 	fmt.Println(featuredPrice)
// 	fmt.Println(highlightedPrice)
// }

// Time to practice what you learned!

type Products struct {
	id    string
	title string
	price float64
}

func main() {
	var hobbies [3]string = [3]string{"Eat", "Sleep", "Code"}
	fmt.Println(hobbies[:1])
	fmt.Println(hobbies[1:3])

	var courseGoals []string = []string{"Learn Go", "Build Projects"}
	fmt.Println(courseGoals)
	courseGoals[1] = "Chalo"
	fmt.Println(courseGoals)
	courseGoals = append(courseGoals, "Become a Go Expert")
	fmt.Println(courseGoals)

	var products []Products = []Products{
		{title: "Fan", id: "1", price: 10.99},
		{title: "Light", id: "2", price: 19.99},
	}
	products[0] = Products{title: "Light1", id: "2", price: 19.99}
	products[1].price = 29.99
	products = append(products, Products{title: "Mobile", id: "2", price: 19.99})
	fmt.Println(products)
}

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
