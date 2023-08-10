package interestcalculation

import (
	"fmt"
	"math"
)

// Exercise 8: Interest Calculation
// Create a function that calculates the final value of an investment based on initial capital,
// interest rate, and investment time (in months). The program must prompt the user for these
// values and display the final value.
func InterestCalculation(initialCapital, interestRate float64, months int) float64 {
	return initialCapital * math.Pow(1+interestRate, float64(months))
}

func Run() {
	fmt.Println("Please provide the initial capital:")
	var initialCapital float64
	fmt.Scan(&initialCapital)

	fmt.Println("Please provide the interest rate:")
	var interestRate float64
	fmt.Scan(&interestRate)

	fmt.Println("Please provide how many months were invested:")
	var months int
	fmt.Scan(&months)

	interest := InterestCalculation(initialCapital, interestRate, months)
	fmt.Printf("The interest is %f", interest)
}
