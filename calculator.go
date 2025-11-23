package main

import "fmt"

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
