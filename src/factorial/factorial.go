package factorial

import "fmt"

// Exercise 3: Factorial
// Write a function to calculate the factorial of a number. Next, create a program that allows the
// user to enter a number and displays the corresponding factorial.
func Factorial(num int64) int64 {
	if num == 0 {
		return 1
	}

	return num * Factorial(num-1)
}

func Run() {
	chosenOp := 0
	for chosenOp != 2 {
		fmt.Printf("=============================\nPlease choose one of the following options\n\n")
		fmt.Println("1 - Calculate factorial")
		fmt.Println("2 - Exit")
		fmt.Scan(&chosenOp)

		if chosenOp == 1 {
			targetNum := 0
			fmt.Println("Please provide a number to calculate it factorial.")
			fmt.Scan(&targetNum)

			targetFactorial := Factorial(int64(targetNum))
			fmt.Printf("The factorial of %d is %v\n", targetNum, targetFactorial)
		}
	}
}
