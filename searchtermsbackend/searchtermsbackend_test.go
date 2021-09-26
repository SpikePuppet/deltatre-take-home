package main

import (
	"context"
	"log"
	"net"
	"reflect"
	st "searchtermsprotobuf"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const (
	bufsize = 1024 * 1024
)

var listener *bufconn.Listener

func init() {
	listener = bufconn.Listen(bufsize)
	server := grpc.NewServer()
	st.RegisterSearchServer(server, &SearchServer{})

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatalf("server exited with error %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return listener.Dial()
}

func TestSearchTerms(t *testing.T) {
	tables := []struct {
		searchTermRequest st.SearchTermRequest
		result            string
	}{
		{st.SearchTermRequest{Term: "hello"}, "Term exists!"},
		{st.SearchTermRequest{Term: "goodbye"}, "Term exists!"},
		{st.SearchTermRequest{Term: "simple"}, "Term exists!"},
		{st.SearchTermRequest{Term: "list"}, "Term exists!"},
		{st.SearchTermRequest{Term: "search"}, "Term exists!"},
		{st.SearchTermRequest{Term: "filter"}, "Term exists!"},
		{st.SearchTermRequest{Term: "yes"}, "Term exists!"},
		{st.SearchTermRequest{Term: "no"}, "Term exists!"},
		{st.SearchTermRequest{Term: "seed"}, "Term does not exist!"},
		{st.SearchTermRequest{Term: "database"}, "Term does not exist!"},
		{st.SearchTermRequest{Term: "golang"}, "Term does not exist!"},
		{st.SearchTermRequest{Term: "deltatre"}, "Term does not exist!"},
		{st.SearchTermRequest{Term: "streaming"}, "Term does not exist!"},
	}

	ctx := context.Background()
	connection, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial bufnet: %v", err)
	}
	defer connection.Close()
	client := st.NewSearchClient(connection)

	for _, table := range tables {
		result, _ := client.SearchTerms(ctx, &table.searchTermRequest)
		if result.Message != table.result {
			t.Errorf("Search result was incorrect for %s, got %s, expected %s", table.searchTermRequest.Term, result.Message, table.result)
		}
	}
}

func TestUpdateTerms(t *testing.T) {
	tables := []struct {
		updateSearchTermRequest string
		result                  string
	}{
		{"seed", "Term exists!"},
		{"database", "Term exists!"},
		{"golang", "Term exists!"},
		{"deltatre", "Term exists!"},
		{"streaming", "Term exists!"},
	}

	ctx := context.Background()
	connection, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial bufnet: %v", err)
	}
	defer connection.Close()
	client := st.NewSearchClient(connection)

	for _, table := range tables {
		_, err = client.UpdateTerms(ctx, &st.UpdateSearchTermsRequest{Term: table.updateSearchTermRequest})
		if err != nil {
			t.Errorf("Test failed, an error occured during update, err %v", err)
		}

		result, _ := client.SearchTerms(ctx, &st.SearchTermRequest{Term: table.updateSearchTermRequest})
		if result.Message != table.result {
			t.Errorf("Test failed, search result was incorrect for %s, got %s, expected %s", table.updateSearchTermRequest, result.Message, table.result)
		}
	}
}

func TestGetSearchMetrics(t *testing.T) {
	tables := []struct {
		searchTerms     []string
		expectedMetrics []*st.SearchTermMetrics
	}{
		{[]string{"hello", "hello", "hello", "simple", "simple", "yes"}, []*st.SearchTermMetrics{
			{Term: "hello", SearchCount: 3},
			{Term: "simple", SearchCount: 2},
			{Term: "yes", SearchCount: 1},
		}},
		{[]string{"hello", "goodbye", "goodbye", "goodbye", "goodbye", "goodbye", "simple", "simple", "simple", "simple", "yes", "yes", "yes", "no", "no"},
			[]*st.SearchTermMetrics{
				{Term: "goodbye", SearchCount: 5},
				{Term: "simple", SearchCount: 4},
				{Term: "yes", SearchCount: 3},
				{Term: "no", SearchCount: 2},
				{Term: "hello", SearchCount: 1},
			},
		},
	}

	ctx := context.Background()
	connection, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial bufnet: %v", err)
	}
	defer connection.Close()
	client := st.NewSearchClient(connection)

	for _, table := range tables {
		terms = map[string]int{
			"hello":   0,
			"goodbye": 0,
			"simple":  0,
			"list":    0,
			"search":  0,
			"filter":  0,
			"yes":     0,
			"no":      0,
		}

		for _, value := range table.searchTerms {
			_, err = client.SearchTerms(ctx, &st.SearchTermRequest{Term: value})
			if err != nil {
				t.Errorf("error during test setup, err: %v", err)
			}
		}

		result, _ := client.GetSearchMetrics(ctx, &st.SearchTermMetricsRequest{})
		for i := 0; i < len(table.expectedMetrics)-1; i++ {
			if !reflect.DeepEqual(table.expectedMetrics[i], result.Results[i]) {
				t.Errorf("Test failed, metrics order was incorrect, got %v, expected %v", result.Results[i], table.expectedMetrics[i])
			}
		}
	}
}
