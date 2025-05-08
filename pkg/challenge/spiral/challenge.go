package spiral

import (
	"fmt"
	"math"
)

const (
	Left int = iota
	Up
	Right
	Down
)

func generateSpiralMatrix(n int) [][]string {
	size := int(math.Round(math.Sqrt(float64(n + 1))))
	if n%size != 0 {
		size++
	}

	matrix := make([][]string, size)
	for i := range matrix {
		matrix[i] = make([]string, size)
	}

	x, y := size/2, size/2
	direction := Left
	step, step_limit := 0, 1

	for i := range n + 1 {
		value := fmt.Sprintf("%3d", i)
		matrix[y][x] = value

		if direction == Left {
			x--
		} else if direction == Up {
			y--
		} else if direction == Right {
			x++
		} else if direction == Down {
			y++
		}

		step++
		if step == step_limit {
			direction = (direction + 1) % 4

			if direction == Left || direction == Right {
				step_limit++
			}
			step = 0
		}
	}

	return matrix
}

func printMatrix(matrix [][]string) {
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%s ", val)
		}
		fmt.Println()
	}
}
