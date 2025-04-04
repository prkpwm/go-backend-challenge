
package main

import (
	"fmt"
	"math"
	"strings"
)

// Memoization to store results for (index, previous digit)
var memo map[string][]int

// Generate the minimum sum sequence using recursion and memoization
func findMinSequence(encoded string, index, prevDigit, n int) []int {
	// Base case: reached the end, return empty sequence
	if index == n {
		return []int{}
	}

	// Create a unique key for memoization
	key := fmt.Sprintf("%d-%d", index, prevDigit)
	if val, found := memo[key]; found {
		return val
	}

	// Store the minimum sequence found
	minSum := math.MaxInt64
	var minSeq []int

	// Try all digits from 0 to 9
	for digit := 0; digit <= 9; digit++ {
		// Ensure the digit follows the encoded constraints
		if index == 0 || 
		   (encoded[index-1] == 'L' && prevDigit > digit) ||
		   (encoded[index-1] == 'R' && prevDigit < digit) ||
		   (encoded[index-1] == '=' && prevDigit == digit) {

			// Recursively find the next best sequence
			nextSeq := findMinSequence(encoded, index+1, digit, n)

			// Calculate the sum of this sequence
			totalSum := digit
			for _, num := range nextSeq {
				totalSum += num
			}

			// Update if we found a smaller sum sequence
			if totalSum < minSum {
				minSum = totalSum
				minSeq = append([]int{digit}, nextSeq...)
			}
		}
	}

	// Store in memoization and return
	memo[key] = minSeq
	return minSeq
}

// Decode the encoded string into the smallest valid number
func decode(encoded string) string {
	memo = make(map[string][]int) // Reset memoization
	n := len(encoded) + 1

	// Find the smallest valid sequence
	minSeq := findMinSequence(encoded, 0, -1, n)

	// Convert the sequence to a string
	var result strings.Builder
	for _, num := range minSeq {
		result.WriteString(fmt.Sprintf("%d", num))
	}

	return result.String()
}

func main() {
	var encoded string
	fmt.Print("input = ")
	fmt.Scan(&encoded)

	decoded := decode(encoded)
	fmt.Printf("output = %s\n", decoded)
}