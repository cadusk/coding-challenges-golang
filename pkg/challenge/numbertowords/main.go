package numbertowords

import (
	"fmt"
	"strconv"
)

func RunChallenge() {
	fmt.Println("Enter an amount:")

	amounts := []float64{2523.15, 13.763, 13.768, 511.0, 0.0, 0.5}
	for _, amount := range amounts {
		fmt.Printf("The amount to convert is '%.2f' and that in words is: %s\n", amount, numberToWords(amount))
	}

	fmt.Println("") // make some room

	stringAmounts := []string{"2523.15", "13.763", "13.768", "511.0", "0.0", "0.5", "blah", "42.5 "}
	for _, stringAmount := range stringAmounts {
		amount, err := strconv.ParseFloat(stringAmount, 64)
		if err != nil {
			fmt.Printf("Error converting string to float: %v\n", err)
			continue
		}

		fmt.Printf("The amount to convert is '%.2f' and that in words is: %s\n", amount, numberToWords(amount))
	}
}
