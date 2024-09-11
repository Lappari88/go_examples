package main

import "fmt"

// Basic functions. Variable type is after the parameter name and return value is last

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}
// Multiple retrun values
func math_action(a in, b int, action string)(int,string) {
	if action == "add" {
		return a + b, "added"
	}
	if action == "subtract" {
		return a - b,"subtracted"
	}
	return 0, "unknown action"
}

//Variadic functions

func sum(vals..int) int {
	total := 0
	for _,num := range vals {
		total += num
	}
	return total
}

func main() {
	c := add(20, 10)
	d := subtract(20, 10)
	fmt.Println("20 + 10 =", c)
	fmt.Println("20 - 10 =", d)

	e,act := math_action(20,10,"add")
	f,act2 := math_action(20,10,"subtract")
}
