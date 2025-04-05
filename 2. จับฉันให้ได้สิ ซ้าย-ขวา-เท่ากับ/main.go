package main

import (
    "fmt"
    "math"
    "strings"
)

func main() {
    var encoded string
    fmt.Print("input = ")
    fmt.Scan(&encoded)

    decoded := decode(encoded)
    fmt.Printf("output = %s\n", decoded)
}

func decode(encoded string) string {
    n := len(encoded) + 1
    results := generatePossibleNumbers(encoded, n)
    minSequence := findMinSumSequence(results)

    var builder strings.Builder
    for _, num := range minSequence {
        builder.WriteString(fmt.Sprintf("%d", num))
    }

    return builder.String()
}

func generatePossibleNumbers(encoded string, n int) [][]int {
    results := make([][]int, 0)
    generateNumbers(encoded, n, 0, []int{}, &results)
    return results
}

func findMinSumSequence(results [][]int) []int {
    minSum := math.MaxInt64
    var minSequence []int

    for _, seq := range results {
        sum := 0
        for _, num := range seq {
            sum += num
        }
        if sum < minSum {
            minSum = sum
            minSequence = seq
        }
    }

    return minSequence
}

func generateNumbers(encoded string, n, index int, current []int, results *[][]int) {
    if index == n {
        *results = append(*results, append([]int(nil), current...))
        return
    }
    for i := 0; i <= 9; i++ {
        if isValid(encoded, index, current, i) {
            generateNumbers(encoded, n, index+1, append(current, i), results)
        }
    }
}

func isValid(encoded string, index int, current []int, next int) bool {
    if index == 0 {
        return true
    }
    prev := current[index-1]
    switch encoded[index-1] {
    case 'L':
        return prev > next
    case 'R':
        return prev < next
    case '=':
        return prev == next
    }
    return false
}