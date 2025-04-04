package main

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "regexp"
    "strings"

    "github.com/gorilla/mux"
)

type MeatSummary struct {
    Beef map[string]int `json:"beef"`
}

var sourceURL string = "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"

func fetchText() (string, error) {
    resp, err := http.Get(sourceURL)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }
    return string(body), nil
}

func countMeats(text string) map[string]int {
    re := regexp.MustCompile(`[^\w\s]`)
    cleanText := re.ReplaceAllString(text, "")
    words := strings.Fields(cleanText)
    counts := make(map[string]int)
    for _, word := range words {
        counts[strings.ToLower(word)]++
    }
    return counts
}

func beefSummaryHandler(w http.ResponseWriter, r *http.Request) {
    text, err := fetchText()
    if err != nil {
        http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
        return
    }
    response := MeatSummary{Beef: countMeats(text)}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    response := map[string]string{
        "error":   "Path not found",
        "message": "Available paths: /beef/summary",
    }
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode(response)
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/beef/summary", beefSummaryHandler).Methods("GET")
    router.NotFoundHandler = http.HandlerFunc(notFoundHandler)
    fmt.Println("Server running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
