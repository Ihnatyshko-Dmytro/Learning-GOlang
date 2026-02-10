package main

import "fmt"

func main() {
	defer hendlerPanic()

	messages := []string {
		"message 1",
		"message 2",
		"message 3",
		"message 4",
	}
	fmt.Println("hendlerPanic is active")
	messages[4] = "message 5"

	fmt.Println(messages)
	// fmt.Println("main()")
}

// func printMessage() {
// 	fmt.Println("printMessage")
// }
func hendlerPanic() {
	if r := recover(); r!= nil {
		fmt.Println(r)
	}
	fmt.Println("hendlerPanic() to do succsesfully")
}