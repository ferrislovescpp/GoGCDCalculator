package main

import "fmt"

// calculateGCD calculates the greatest common divisor of two integers using the Euclidean algorithm
// we do this to find the largest positive integer that divides both "a" and "b" without a remainder
func calculateGCD(a int, b int) int {
	// the algorithm continues until b becomes zero
	// at which point a will contain the GCD, using a while-loop that checks if b is zero
	for b != 0 {
		// swap a and b, and set b to a mod b to apply the euclidean algo ( GCD(a, b) = GCD(b, a mod b) )
		// we do this through using Go's multiple assignment feature to swap a and b, and calculate a mod b
		a, b = b, a%b
	}

	// here we just return "a" because it contains the GCD according to the Euclidean algorithm
	return a
}

// isValidChoice checks if the user's choice is valid by checking if the input string matches either "y" or "n"
// to ensure the user has entered a recognizable command before proceeding.
func isValidChoice(choice string) bool {
	// using logical OR to check if choice is either "y" or "n" compare choice against allowed options "y" and "n"
	// to validate that the user's input is one of the expected choices for proceeding.
	return choice == "y" || choice == "n"
}

// this is the entry point of every program, "main"
func main() {
	for {
		// using Go's var keyword to declare variables, we declare variables for user choice and numbers
		// to store the user's input for further processing
		var choice string
		var num1, num2 int

		// using fmt.Print to display a message without a newlin to prompt the user for their choice
		// to ask the user if they want to proceed with GCD calculation.
		fmt.Print("do you want to calculate the gcd of two non-negative integers? (y/n) ")

		// using fmt.Scanln to read a line of input from the standard input
		// read the user's choice to know whether to proceed with the rest of the program or not.
		fmt.Scanln(&choice)

		// validate the choice by calling the isValidChoice function and checking its return value
		/// to ensure the program proceeds only if a valid choice is made.
		if !isValidChoice(choice) {
			fmt.Println("invalid choice. exiting.")
			return
		}

		// by checking if choice equals "n" and using return to exit the program
		// exit if the choice is 'n' because the user chose not to proceed.
		if choice == "n" {
			fmt.Println("exiting.")
			return
		}

		// prompt the user to enter two numbers to get the numbers for which the GCD will be calculated
		// using fmt.Print to display a message without a newline.
		fmt.Print("enter two non-negative integers: ")
		fmt.Scanln(&num1, &num2)

		// read the numbers from the user to get the actual integers for GCD calculation
		// using fmt.Scanln to read the integers from the standard input
		if num1 < 0 || num2 < 0 {
			fmt.Println("please enter non-negative integers. exiting.")
			return
		}

		// calculate GCD and display the result
		gcd := calculateGCD(num1, num2)
		fmt.Printf("the gcd of %d and %d is %d\n", num1, num2, gcd)
	}
}

