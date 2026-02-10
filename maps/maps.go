package main

import "fmt"

func main() {
	users := map[string]int{
		"Vasya":  15,
		"Petya":  23,
		"Kostya": 45,
	}

	age, exists := users["Kostya"]

	fmt.Println(exists)

	if exists {
		fmt.Println("Kostya", age)
	} else {
		fmt.Println("not on the list")
	}
}
