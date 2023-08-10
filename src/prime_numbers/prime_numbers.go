package primenumbers

import (
	"fmt"
	"math"
)

// Exercise 2: Prime Numbers
// Write a function that checks whether a number is prime or not. Then create a program that
// prints the first 10 prime numbers.
func PrimeNumbers(num int) bool {
	if num <= 1 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

// FirstTenPrimes is the helper function requested at exercise 2.
func FirstTenPrimes() {
	fmt.Println("Testing the exercise 2")
	i := 0
	primeCounter := 0

	for primeCounter < 10 {
		if PrimeNumbers(i) {
			fmt.Printf("%d is prime\n", i)
			primeCounter++
		}
		i++
	}
}
