package main

import (
    "fmt"
    "math"
    "strings"
)

func generatePossibleNumbers(encoded string, n int) [][]int {
    results := make([][]int, 0)

    var generate func(index int, current []int)
    generate = func(index int, current []int) {
        if index == n {
            results = append(results, append([]int(nil), current...))
            return
        }
        for i := 0; i <= 9; i++ {
            if index == 0 || (encoded[index-1] == 'L' && current[index-1] > i) ||
                (encoded[index-1] == 'R' && current[index-1] < i) ||
                (encoded[index-1] == '=' && current[index-1] == i) {
                generate(index+1, append(current, i))
            }
        }
    }

    generate(0, []int{})
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

func main() {
    var encoded string
    fmt.Print("input = ")
    fmt.Scan(&encoded)

    decoded := decode(encoded)
    fmt.Printf("output = %s\n", decoded)
}
