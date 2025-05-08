package pokerrank_test

import (
	"testing"

	"programmersinc.com/challenges/pkg/challenge/pokerrank"
)

func TestSortingBeforeRanking(t *testing.T) {
	hand := []string{"2H", "3D", "5S", "9C", "KD"}

	expected := []string{"KD", "9C", "5S", "3D", "2H"}
	result, _ := pokerrank.NewHand(hand)

	for i, card := range result {
		if card.String() != expected[i] {
			t.Errorf("Expected %s, got %s", expected[i], card.String())
		}
	}
}
