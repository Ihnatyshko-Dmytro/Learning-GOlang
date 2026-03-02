package main

import "fmt"

type Users struct {
		email string
		name string
	}

func main() {

	a := []int64{1, 2, 3}
	b := []float64{1.1, 2.2, 3.3}
	c := []string{"1", "2", "3"}

	d := []Users {
		{
			email: "dfsdf@gmail.com",
			name: "Alex",
		},
		{
			email: "df1f@gmail.com",
			name: "Tolik",
		},
		{
			email: "dfsdf@gmail.com",
			name: "Misha",
		},
	}

	// fmt.Println(sumOfFloat64(b))
	// fmt.Println(sumOfInt64(a))
	fmt.Println(sum(b))
	fmt.Println(sum(a))
	fmt.Println(searchElement(c, "2"))
	fmt.Println(searchElement(d, Users{
		email: "dfsdf@gmail.com",
		name: "Misha",
	}))

}

func sum[V int64 | float64](input []V) V {
	var result V
	for _, number := range input {
		result += number
	}
	return result
}

func searchElement[C comparable](elements []C, searchEl C) bool {
	for _, el := range elements {
		if el == searchEl {
			return true
		}
	}
	return false
}

// func sumOfFloat64(input []float64) float64 {
// 	var result float64

// 	for _, number := range input {
// 		result += number
// 	}

// 	return result
// }
