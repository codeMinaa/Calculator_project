package main

import "fmt"

func getSecondNumber() int {
	var b int
	fmt.Print("Enter Second Number:.")
	_, err = fmt.Scanln(&b)
	if err == nil {
		fmt.Print(b)
	} else {
		fmt.Println("Invalid Input!")

	}
	return b
}
