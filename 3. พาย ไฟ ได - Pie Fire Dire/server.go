package main

import (
    "net/http"
    "io"
    "context"
    "log"
    "net"
    "regexp"
    "strings"

    "google.golang.org/grpc"
    pb "go-grpc/proto"
)

type server struct {
    pb.UnimplementedMeatServiceServer
}

var sourceURL string = "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"


func (s *server) GetBeefSummary(ctx context.Context, in *pb.Empty) (*pb.MeatSummary, error) {
    text, err := fetchText()
    if err != nil {
        return nil, err
    }
    beefCount := countMeats(text)
    summary := &pb.MeatSummary{Beef: beefCount}
    return summary, nil
}

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

func countMeats(text string) map[string]int32 {
    re := regexp.MustCompile(`[^\w\s]`)
    cleanText := re.ReplaceAllString(text, "")
    words := strings.Fields(cleanText)
    counts := make(map[string]int32)
    for _, word := range words {
        counts[strings.ToLower(word)]++
    }
    return counts
}

func main() {
    lis, err := net.Listen("tcp", ":8080")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    grpcServer := grpc.NewServer()
    pb.RegisterMeatServiceServer(grpcServer, &server{})
    log.Println("gRPC server running on port 8080...")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
