package simplecalculator

import "fmt"

// Exercise 1: Simple Calculator
// Create a calculator that takes two numbers and an operator (+, -, *, /) and returns the result
// of the operation.
func SimpleCalculator(x, y int, op string) (int, error) {
	switch op {
	case "+":
		return x + y, nil
	case "-":
		return x - y, nil
	case "*":
		return x * y, nil
	case "/":
		if y == 0 {
			return 0, fmt.Errorf("the divisor can not be equal to 0")
		}
		return x / y, nil
	default:
		return 0, fmt.Errorf("the operator %q is not valid", op)
	}
}

func Run() {
	fmt.Println("Please provide a number:")
	var x int
	fmt.Scan(&x)

	fmt.Println("Please an operation:")
	var op string
	fmt.Scan(&op)

	fmt.Println("Please provide the second number:")
	var y int
	fmt.Scan(&y)

	if result, err := SimpleCalculator(x, y, op); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The result of this operation is %d\n", result)
	}
}
