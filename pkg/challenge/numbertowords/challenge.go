package numbertowords

import (
	"math"
	"strconv"
)

// write a program that takes an amount as input and prints the amount in words
// for example, if the input is 1234, the output should be "one thousand two hundred thirty four"
func numberToWords(amount float64) string {
	// Define the words for numbers
	ones := []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	teens := []string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	tens := []string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

	// convert the amount to an integer and a decimal
	whole, decimal := math.Modf(amount)

	num := int(whole)
	cents := int((decimal * 100) + 0.5) // round to the nearest cent

	if num == 0 && cents == 0 {
		return "zero"
	}

	result := ""

	if num >= 1000 {
		result += ones[num/1000] + " thousand "
		num %= 1000
	}

	if num >= 100 {
		result += ones[num/100] + " hundred "
		num %= 100
	}

	if num >= 20 {
		result += tens[num/10] + " "
		num %= 10
	}

	if num >= 10 {
		result += teens[num-10] + " "
		num = 0
	}

	if num > 0 {
		result += ones[num] + " "
	}

	if cents > 0 {
		if result != "" {
			result += "and "
		}
		result += strconv.FormatInt(int64(cents), 10) + "/100 "
	}

	return result[:len(result)-1]
}
