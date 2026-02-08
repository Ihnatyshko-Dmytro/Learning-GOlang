package main

import "fmt"

func main() {

	matrix := make([][]int, 10)
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			matrix[y] = make([]int, 10)
			matrix[y][x] = x

		}
		fmt.Println(matrix)
	}
}
