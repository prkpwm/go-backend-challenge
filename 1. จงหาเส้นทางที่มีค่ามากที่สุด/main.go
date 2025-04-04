package main

import (
    "encoding/json"
    "fmt"
    "os"
)

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
    pathSumMatrix := make([][]int, n)

    for i := range pathSumMatrix {
        pathSumMatrix[i] = make([]int, len(triangle[i]))
    }

    for i := 0; i < len(triangle[n-1]); i++ {
        pathSumMatrix[n-1][i] = triangle[n-1][i]
    }

    for i := n - 2; i >= 0; i-- {
        for j := 0; j < len(triangle[i]); j++ {
            pathSumMatrix[i][j] = triangle[i][j] + max(pathSumMatrix[i+1][j], pathSumMatrix[i+1][j+1])
        }
    }

    return pathSumMatrix[0][0]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    filePath := "hard.json"

    triangle, err := readTriangleFromFile(filePath)
    if err != nil {
        fmt.Println(err)
        return
    }

    maxSum := calculateMaxPathSum(triangle)
    fmt.Println("ค่าของเส้นทางที่มีค่ามากที่สุดคือ:", maxSum)
}
