// write some code that evaluates a poker hand and determines its rank.
package pokerrank

import (
	"fmt"
)

type Card struct {
	Rank int
	Suit string
	Rep  string
}

func (c Card) String() string {
	return c.Rep
}

func NewHand(c []string) ([]Card, error) {
	if len(c) != 5 {
		return nil, fmt.Errorf("Poker hand should always be 5. Got %d\n", len(c))
	}

	hand := []Card{
		parseCard(c[0]),
		parseCard(c[1]),
		parseCard(c[2]),
		parseCard(c[3]),
		parseCard(c[4]),
	}

	// sort cards by rank, highest to lowest
	for i := 0; i < len(hand)-1; i++ {
		for j := 0; j < len(hand)-i-1; j++ {
			if hand[j].Rank < hand[j+1].Rank {
				hand[j], hand[j+1] = hand[j+1], hand[j]
			}
		}
	}

	return hand, nil
}

func parseCard(card string) Card {
	// Map card ranks to integers
	rankMap := map[string]int{
		"2": 2, "3": 3, "4": 4, "5": 5,
		"6": 6, "7": 7, "8": 8, "9": 9,
		"T": 10, "10": 10,
		"J": 11, "Q": 12, "K": 13, "A": 14,
	}

	return Card{
		Rank: func() int {
			if len(card) > 2 && rankMap[card[:2]] != 0 {
				return rankMap[card[:2]]
			}
			return rankMap[card[:1]]
		}(),
		Suit: string(card[len(card)-1]),
		Rep:  card,
	}
}

func rank(hand []Card) string {
	if len(hand) != 5 {
		return "Invalid hand"
	}

	// Check for flush
	isFlush := true
	for i := 1; i < len(hand); i++ {
		if hand[i].Suit != hand[0].Suit {
			isFlush = false
			break
		}
	}

	// Check for straight
	isStraight := true
	for i := 1; i < len(hand); i++ {
		if hand[i].Rank != hand[i-1].Rank-1 {
			isStraight = false
			break
		}
	}

	// Check for pairs, three of a kind, four of a kind
	rankCount := make(map[int]int)
	for _, card := range hand {
		rankCount[card.Rank]++
	}

	// Check for pairs, three of a kind, and four of a kind
	pairs := 0
	threeOfAKind := 0
	fourOfAKind := 0
	for _, count := range rankCount {
		if count == 2 {
			pairs++
		} else if count == 3 {
			threeOfAKind++
		} else if count == 4 {
			fourOfAKind++
		}
	}

	// Determine the rank
	if isFlush && isStraight && hand[0].Rank == 14 {
		return "Royal Flush"
	} else if isFlush && isStraight {
		return "Straight Flush"
	} else if fourOfAKind == 1 {
		return "Four of a Kind"
	} else if threeOfAKind == 1 && pairs == 1 {
		return "Full House"
	} else if isFlush {
		return "Flush"
	} else if isStraight {
		return "Straight"
	} else if threeOfAKind == 1 {
		return "Three of a Kind"
	} else if pairs == 2 {
		return "Two Pair"
	} else if pairs == 1 {
		return "One Pair"
	}

	return "High Card"
}

func RankHand(hand []string) (string, error) {
	pokerHand, err := NewHand(hand)
	if err != nil {
		return "", err
	}
	return rank(pokerHand), nil
}
