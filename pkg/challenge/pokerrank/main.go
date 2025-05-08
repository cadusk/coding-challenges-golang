package pokerrank

import "fmt"

func RunChallenge() {
	hands := [][]string{
		{"AH", "KH", "QH", "JH", "TH"},       // Royal Flush
		{"TH", "9H", "8H", "7H", "6H"},       // Straight Flush
		{"AH", "AD", "AS", "AC", "KH"},       // Four of a Kind
		{"AH", "AD", "AS", "KS", "KH"},       // Full house
		{"KC", "TC", "8C", "7C", "5C"},       // Flush
		{"TH", "9C", "6D", "7C", "8H"},       // Straight
		{"AH", "AD", "AC", "KS", "QH"},       // Three of a Kind
		{"7H", "AC", "KS", "KD", "AH"},       // Two Pair
		{"AH", "AC", "KD", "JS", "7H"},       // One Pair
		{"AH", "3D", "4S", "5C", "9H"},       // High Card
		{"3D", "4S", "5C", "9H"},             // Invalid hand
		{"AC", "AH", "3D", "4S", "5C", "9H"}, // Invalid hand
	}

	for _, h := range hands {
		// Example poker hand
		hand, err := NewHand(h)
		if err != nil {
			fmt.Println("Error creating hand:", err)
			continue
		}

		// Print the rank
		fmt.Println("The hand", hand, "is a", rank(hand))
	}
}
