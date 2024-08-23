package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func divide(a float64, b float64) (result float64, err error) {
	if b == 0 {
		err = fmt.Errorf("Division by zero")
		return
	}
	result = a / b
	return
}

func main() {
	fmt.Println(add(2, 3))
}
