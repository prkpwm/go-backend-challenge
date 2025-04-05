package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	filePath := "triangle.json"

	triangle, err := readTriangleFromFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	maxSum := calculateMaxPathSum(triangle)
	fmt.Println("ค่าของเส้นทางที่มีค่ามากที่สุดคือ:", maxSum)
}

func readTriangleFromFile(filePath string) ([][]int, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	var triangle [][]int
	err = json.Unmarshal(data, &triangle)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %v", err)
	}
	return triangle, nil
}

func calculateMaxPathSum(triangle [][]int) int {
	n := len(triangle)

	for i := n - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += max(triangle[i+1][j], triangle[i+1][j+1])
		}
	}

	maxPathSum := triangle[0][0]
	return maxPathSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
