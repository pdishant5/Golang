package main

import "fmt"

type Product struct {
	ID    int
	title string
	price float64
}

func main() {
	// 1
	hobbies := [3]string{"Cricket", "Traveling", "Coding"}
	fmt.Println("My Hobbies:", hobbies)

	// 2
	fmt.Println("My first hobby is:", hobbies[0])
	fmt.Println("My other hobbies are:", hobbies[1:])

	// 3
	// partialHobbies := hobbies[0:2]
	partialHobbies := hobbies[:2]
	fmt.Println("My partial hobbies are:", partialHobbies)
	fmt.Println("Length:", len(partialHobbies), "Capacity:", cap(partialHobbies)) // 2, 3

	// 4
	partialHobbies = partialHobbies[1:cap(partialHobbies)]
	fmt.Println("My updated partial hobbies are:", partialHobbies)
	fmt.Println("Length:", len(partialHobbies), "Capacity:", cap(partialHobbies)) // 2, 2

	// 5
	courseGoals := []string{"Learn Go Basics", "Learn Go Internals"}
	fmt.Println("Course Goals:", courseGoals)
	fmt.Println("Length:", len(courseGoals), "Capacity:", cap(courseGoals)) // 2, 2

	// 6
	courseGoals[1] = "Learn Go Advanced"
	courseGoals = append(courseGoals, "Build Go Projects")
	fmt.Println("Updated Course Goals:", courseGoals)
	fmt.Println("Length:", len(courseGoals), "Capacity:", cap(courseGoals)) // 3, 4

	// 7
	products := make([]Product, 0)
	products = append(products, Product{ID: 1, title: "Jeans", price: 999.99})
	products = append(products, Product{ID: 2, title: "T-shirt", price: 499.99})
	fmt.Println("Products:", products)
	fmt.Println("Length:", len(products), "Capacity:", cap(products)) // 2, 2

	products = append(products, Product{ID: 3, title: "Sneakers", price: 1999.99})
	fmt.Println("Updated Products:", products)
	fmt.Println("Length:", len(products), "Capacity:", cap(products)) // 3, 4

	// products = products[1:cap(products)]
	// fmt.Println("Updated Products:", products) // products[2] - nil value of "Product" struct
}
