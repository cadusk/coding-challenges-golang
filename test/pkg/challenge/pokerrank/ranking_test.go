package pokerrank_test

import (
	"testing"

	"programmersinc.com/challenges/pkg/challenge/pokerrank"
)

func TestPokerHandRanking_RoyalFlush(t *testing.T) {
	hand := []string{"AH", "KH", "QH", "JH", "TH"}

	expected := "Royal Flush"
	result, _ := pokerrank.RankHand(hand)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestPokerHandRanking_StraightFlush(t *testing.T) {
	hand := []string{"TH", "9H", "8H", "7H", "6H"}

	expected := "Straight Flush"
	result, _ := pokerrank.RankHand(hand)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestPokerHankRanking_FourOfAKind(t *testing.T) {
	hand := []string{"AH", "AD", "AS", "AC", "KH"}

	expected := "Four of a Kind"
	result, _ := pokerrank.RankHand(hand)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestPokerHandRanking_FullHouse(t *testing.T) {
	hand := []string{"AH", "AD", "AS", "KS", "KH"}

	expected := "Full House"
	result, _ := pokerrank.RankHand(hand)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestPokerHandRanking_Flush(t *testing.T) {
	hand := []string{"AH", "KH", "QH", "JH", "5H"}

	expected := "Flush"
	result, _ := pokerrank.RankHand(hand)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestPokerHandRanking_Straight(t *testing.T) {
	hand := []string{"TH", "9D", "8S", "7C", "6H"}

	expected := "Straight"
	result, _ := pokerrank.RankHand(hand)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestPokerHandRanking_ThreeOfAKind(t *testing.T) {
	hand := []string{"AH", "AD", "AS", "KH", "5H"}

	expected := "Three of a Kind"
	result, _ := pokerrank.RankHand(hand)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestPokerHandRanking_TwoPair(t *testing.T) {
	hand := []string{"AH", "AD", "KH", "KS", "5H"}

	expected := "Two Pair"
	result, _ := pokerrank.RankHand(hand)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestPokerHandRanking_OnePair(t *testing.T) {
	hand := []string{"AH", "AD", "KH", "5S", "3H"}

	expected := "One Pair"
	result, _ := pokerrank.RankHand(hand)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestPokerHandRanking_HighCard(t *testing.T) {
	hand := []string{"AH", "KD", "QS", "JH", "5C"}

	expected := "High Card"
	result, _ := pokerrank.RankHand(hand)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
