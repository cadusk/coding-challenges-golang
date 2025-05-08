package spiral

import "fmt"

func RunChallenge() {

	var input int
	fmt.Println("Enter a number:")
	fmt.Scan(&input)

	if input < 2 {
		fmt.Println("Please enter a non-negative integer greater than 1.")
		return
	}

	matrix := generateSpiralMatrix(input)
	printMatrix(matrix)
}
