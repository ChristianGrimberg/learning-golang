package main

import (
	"fmt"
)

func main() {
	const initial string = "Hello"
	var name string = "world"
	var num1 int = 10
	var num2 int = 5
	var cup1 Coffee
	var cup2 Coffee

	cup1 = Coffee{name: "Express", price: 5, sugar: true, milk: 250}
	cup2 = Coffee{"Milk", 7, true, 350}

	fmt.Printf("%s %s!\n", initial, name)
	fmt.Printf("\nResult: %d\n", num1*num2)
	fmt.Printf("\n==Coffee==\nName: %s\nPrice: %.2f\nSugar: %t\nMilk: %d\n", cup1.name, cup1.price, cup1.sugar, cup1.milk)
	fmt.Printf("\n==Coffee==\nName: %s\nPrice: %.2f\nSugar: %t\nMilk: %d\n", cup2.name, cup2.price, cup2.sugar, cup2.milk)
	fmt.Printf("\nMuliply %d x %d = %d\n", num1, num2, multiply(num1, num2))
}

// Coffee type
type Coffee struct {
	name  string
	price float32
	sugar bool
	milk  int
}

func multiply(num1 int, num2 int) int {
	var result int

	result = num1 * num2

	return result
}
