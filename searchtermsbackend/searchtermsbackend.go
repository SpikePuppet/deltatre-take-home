package main

import (
	"context"
	"log"
	"net"
	"sort"

	"google.golang.org/grpc"
)

const (
	address = "backend:50080"
)

type SearchTermsServer struct {
	UnimplementedSearchServer
}

var terms = map[string]int{
	"hello":   0,
	"goodbye": 0,
	"simple":  0,
	"list":    0,
	"search":  0,
	"filter":  0,
	"yes":     0,
	"no":      0,
}

func main() {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	RegisterSearchServer(server, &SearchTermsServer{})
	log.Printf("server listening at %v", listener.Addr())
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *SearchTermsServer) SearchTerms(ctx context.Context, in *SearchTermRequest) (*SearchTermResponse, error) {
	_, exists := terms[in.Term]
	if !exists {
		return &SearchTermResponse{Message: "Term does not exist!"}, nil
	}

	terms[in.Term]++
	return &SearchTermResponse{Message: "Term exists!"}, nil
}

func (s *SearchTermsServer) UpdateTerms(ctx context.Context, in *UpdateSearchTermsRequest) (*UpdateSearchTermsResponse, error) {
	terms[in.Term] = 0
	return &UpdateSearchTermsResponse{}, nil
}

func (s *SearchTermsServer) GetSearchMetrics(ctx context.Context, in *SearchTermMetricsRequest) (*SearchTermMetricsResponse, error) {
	var results []*SearchTermMetrics
	for key, value := range terms {
		results = append(results, &SearchTermMetrics{Term: key, SearchCount: int32(value)})
	}

	sort.SliceStable(results, func(i, j int) bool {
		return results[i].SearchCount > results[j].SearchCount
	})

	return &SearchTermMetricsResponse{Results: results[:5]}, nil
}
