package main

import "fmt"

func main() {
	// users := map[string]int{
	// 	"Vasya":  15,
	// 	"Petya":  23,
	// 	"Kostya": 45,
	// }

	var users map[string]int
	users = make(map[string]int)

	users["Petya"] = 19
	users["Kostya"] = 44

	delete(users, "Petya")

	age, exists := users["Kostya"]

	fmt.Println(exists)

	if exists {
		fmt.Println("Kostya", age)
	} else {
		fmt.Println("not on the list")
	}

	for key, value := range users {
		 fmt.Println(key, value)
	}
		fmt.Println(len(users))

	first := []int{1, 2, 3}
	second := first 

	first = append(first, 4)

	fmt.Println(&first[0])
	fmt.Println(&second[0])
}
