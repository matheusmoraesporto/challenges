package table

import "fmt"

// Exercise 5: Table
// Create a program that receives a number from the user and displays the table of that
// number, from 1 to 10.
func Table(num int) []int {
	table := make([]int, 10)

	for i := 1; i <= 10; i++ {
		mult := i * num
		table = append(table, mult)
		fmt.Printf("%d x %d = %d\n", num, i, mult)
	}

	return table
}

func Run() {
	fmt.Printf("=============================\nPlease provide a number to show it table:\n\n")
	var num int
	fmt.Scan(&num)
	Table(num)
}
