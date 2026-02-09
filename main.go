package main

import "fmt"

func main() {

	matrix := make([][]int, 10)

	counter := 0

	for x := 0; x < 10; x++ {
		matrix[x] = make([]int, 10)
		for y := 0; y < 10; y++ {
			counter++
			matrix[x][y] = counter

		}
		fmt.Println(matrix[x])
	}
	
	
	
	messages := []string {
		"message 1",
		"message 2",
		"message 3",
		"message 4",
	}
	var copyMess []string
	for i := 0; i < len(messages);  i++ {
		fmt.Println(messages[1])
	}

	for _, message := range messages {
		fmt.Println(message)
		copyMess = append(copyMess, message)
		fmt.Println(copyMess)
	} 

	counter1 := 0 
	for { 
		if counter1 == 100 {
			break
		}
		counter1++ 
		fmt.Println(counter1)
	}

	
}
