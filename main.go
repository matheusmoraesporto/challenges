package main

import (
	"challenge/src/factorial"
	gradeaverage "challenge/src/grade_average"
	interestcalculation "challenge/src/interest_calculation"
	"challenge/src/palindrome"
	primenumbers "challenge/src/prime_numbers"
	simplecalculator "challenge/src/simple_calculator"
	"challenge/src/table"
	"challenge/src/vowel"
	"fmt"
)

func main() {
	for {
		fmt.Println("Choose one of the following exercises:")
		fmt.Println("======================================")
		fmt.Println("1: Simple Calculator")
		fmt.Println("2: Prime Numbers")
		fmt.Println("3: Factorial")
		fmt.Println("4: Palindrome")
		fmt.Println("5: Table")
		fmt.Println("6: Vowel Counter")
		fmt.Println("7: Grade Average")
		fmt.Println("8: Interest Calculation")
		fmt.Println("9: Exit program")
		fmt.Println("======================================")

		chosenOp := 0
		fmt.Scan(&chosenOp)

		switch chosenOp {
		case 1:
			simplecalculator.Run()
		case 2:
			primenumbers.FirstTenPrimes()
		case 3:
			factorial.Run()
		case 4:
			palindrome.Run()
		case 5:
			table.Run()
		case 6:
			vowel.Run()
		case 7:
			gradeaverage.Run()
		case 8:
			interestcalculation.Run()
		case 9:
			return
		default:
			fmt.Println("Invalid option, please try another.")
		}
	}
}
