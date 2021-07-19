package main

import "fmt"

func init() {
	fmt.Printf("==Start with init==\n")
}

func main() {
	const initial string = "Hello" // Variables of constant type
	var name string = "world"      // Default initialization
	var num1, num2 int = 10, 5     // Multiple value assignment
	var num3 = 20                  // Automatic type inference
	price1, price2 := 40, 50.15    // Shorthand assignment
	var cup1, cup2 Coffee

	cup1 = Coffee{name: "Express", price: float32(price1), sugar: true, milk: 250}
	cup2 = Coffee{"Milk", float32(price2), true, 350}

	fmt.Printf("=> %s %s!\n", initial, name)
	fmt.Printf("=> Result: %d\n", num1*num2)
	fmt.Printf("==Coffee==\nName: %s\nPrice: %.2f\nSugar: %t\nMilk: %d\n", cup1.name, cup1.price, cup1.sugar, cup1.milk)
	fmt.Printf("==Coffee==\nName: %s\nPrice: %.2f\nSugar: %t\nMilk: %d\n", cup2.name, cup2.price, cup2.sugar, cup2.milk)
	fmt.Printf("=> Muliply %d x %d = %d\n", num1, num2, multiply(num1, num2))
	value1, value2 := clousures(num3)
	fmt.Printf("=> %d %s\n", value1, value2)
}

// Coffee type
type Coffee struct {
	name  string
	price float32
	sugar bool
	milk  int
}

func multiply(num1 int, num2 int) int {
	var result int = num1 * num2

	return result
}

func clousures(number int) (int, string) {
	variable := func() int {
		return number * 10
	}

	return variable(), "units"
}
