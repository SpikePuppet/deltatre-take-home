package main

import (
	"context"
	"log"
	"net"
	st "searchtermsprotobuf"
	"sort"

	"google.golang.org/grpc"
)

const (
	port = ":50080"
)

type SearchServer struct {
	st.UnimplementedSearchServer
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
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	st.RegisterSearchServer(server, &SearchServer{})
	log.Printf("server listening at %v", listener.Addr())
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *SearchServer) SearchTerms(ctx context.Context, in *st.SearchTermRequest) (*st.SearchTermResponse, error) {
	_, exists := terms[in.Term]
	if !exists {
		return &st.SearchTermResponse{Message: "Term does not exist"}, nil
	}

	terms[in.Term]++
	return &st.SearchTermResponse{Message: "Term exists!"}, nil
}

func (s *SearchServer) UpdateTerms(ctx context.Context, in *st.UpdateSearchTermsRequest) (*st.UpdateSearchTermsResponse, error) {
	terms[in.Term] = 0
	return &st.UpdateSearchTermsResponse{}, nil
}

func (s *SearchServer) GetSearchMetrics(ctx context.Context, in *st.SearchTermMetricsRequest) (*st.SearchTermMetricsResponse, error) {
	var results []*st.SearchTermMetrics
	for key, value := range terms {
		results = append(results, &st.SearchTermMetrics{Term: key, SearchCount: int32(value)})
	}

	sort.SliceStable(results, func(i, j int) bool {
		return results[i].SearchCount > results[j].SearchCount
	})

	return &st.SearchTermMetricsResponse{Results: results[:5]}, nil
}
