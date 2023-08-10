package gradeaverage

import "fmt"

// Exercise 7: Grade Average
// Create a program that takes a student's grades in three different subjects and calculates the
// average of the grades. Then display the calculated average.
func GradeAverage(grades []float32) float32 {
	if len(grades) < 1 {
		return 0
	}

	sum := float32(0)
	for _, grade := range grades {
		sum += grade
	}

	return sum / float32(len(grades))
}

func Run() {
	fmt.Println("Please provide three grades to calculate the average:")

	grades := make([]float32, 3)
	for i := 1; i <= 3; i++ {
		var input float32
		fmt.Printf("Grade %d:\n", i)
		fmt.Scan(&input)
		grades = append(grades, input)
	}

	avg := GradeAverage(grades)
	fmt.Printf("The grade average is %.2f\n", avg)
}
