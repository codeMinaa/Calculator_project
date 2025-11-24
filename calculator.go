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

func getFirstNumber() int {
	var a int
	fmt.Print("Enter First Number:.")
	_, err = fmt.Scanln(&a)
	if err == nil {
		fmt.Println(a)
	} else {
		fmt.Println("Invalid Input!")
	}
	return a
}

func main() {
	fmt.Println("Welcome to CLI Calculator")

	// Call Get First Number function
	firstNumber := getFirstNumber()

	// Call Get Second Number function
	secondNumber := getSecondNumber()

	// Call Get Operation function
	operation := getOperation()

	// Call Perform Calculation function
	result := performCalculation(firstNumber, secondNumber, operation)

	// Display the result to user
	fmt.Println(firstNumber + " " + operation + " " + secondNumber + " = " + result)
}
