package palindrome

import (
	"fmt"
)

func RunChallenge() {
	var input int
	fmt.Println("Enter a number:")
	fmt.Scan(&input)

	if input < 2 {
		fmt.Println("Please enter a non-negative integer greater than 1.")
		return
	}

	if IsPalindrome(input) {
		fmt.Printf("The number '%d' is a palindrome.\n", input)
		return
	}

	fmt.Printf("The number '%d' is not a palindrome.\n", input)
}
