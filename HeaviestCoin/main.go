package main

import (
	"fmt"
)

// a struct to contain the weight and position of the coin
type Coin struct {
	Weight   int
	Position int
}

func findHeavierCoin(coins []Coin) Coin {
	n := len(coins)

	// base case where the last coin is the heaviest
	if n == 1 {
		return coins[0]
	}

	// subparts
	groupSize := n / 3
	groupA := coins[:groupSize]
	groupB := coins[groupSize : 2*groupSize]
	groupC := coins[2*groupSize : 3*groupSize]

	// "weigh" group A and B
	sumA, sumB := 0, 0
	for _, coin := range groupA {
		sumA += coin.Weight
	}
	for _, coin := range groupB {
		sumB += coin.Weight
	}

	// compare sums
	if sumA > sumB {
		// heavier in a
		return findHeavierCoin(groupA)
	} else if sumB > sumA {
		// heavier in b
		return findHeavierCoin(groupB)
	} else {
		// heavier in c
		return findHeavierCoin(groupC)
	}
}

func main() {
	// changing from dynamic input because that takes too long and i am lazy
	// array of coin structs
	coins := []Coin{
		{Weight: 1, Position: 1},
		{Weight: 1, Position: 2},
		{Weight: 1, Position: 3},
		{Weight: 1, Position: 4},
		{Weight: 1, Position: 5},
		{Weight: 1, Position: 6},
		{Weight: 1, Position: 7},
		{Weight: 2, Position: 8},
		{Weight: 1, Position: 9},
	}

	/// call function to find
	heavierCoin := findHeavierCoin(coins)
	fmt.Printf("The heavier coin is at position %d\n", heavierCoin.Position)
}
