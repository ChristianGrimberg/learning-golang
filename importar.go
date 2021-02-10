package main

import (
	"fmt"
)

func main() {
	const initial string = "Hello"
	var name string = "world"
	var num1 int = 10
	var num2 int = 5

	fmt.Printf("%s %s!\nResult: %d\n", initial, name, num1*num2)
}
