package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	st "searchtermsprotobuf"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

const (
	searchtermsbackendaddress = "localhost:50080"
)

func searchTerms(c *gin.Context) {
	conn, err := grpc.Dial(searchtermsbackendaddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("could not connect: %v", err)
		c.JSON(404, gin.H{
			"error": fmt.Sprintf("could not connect: %v", err),
		})
	}
	defer conn.Close()

	client := st.NewSearchClient(conn)

	term := c.Param("term")
	if len(term) <= 0 {
		log.Printf("cannot search with empty search term, %v", err)
		c.JSON(400, gin.H{
			"error": "cannot search with empty search term",
		})
	}

	request := st.SearchTermRequest{
		Term: term,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := client.SearchTerms(ctx, &request)
	if err != nil {
		log.Printf("error updating search terms, err: %v", err)
		c.JSON(400, gin.H{
			"error": fmt.Sprintf("error updating search terms, err: %v", err),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"results": response.Message,
	})
	return

}

func updateSearchTerms(c *gin.Context) {
	conn, err := grpc.Dial(searchtermsbackendaddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("could not connect: %v", err)
		c.JSON(404, gin.H{
			"error": fmt.Sprintf("could not connect: %v", err),
		})
	}
	defer conn.Close()

	client := st.NewSearchClient(conn)

	term := c.Param("term")
	if len(term) <= 0 {
		log.Printf("cannot use empty search term, %v", err)
		c.JSON(400, gin.H{
			"error": "cannot update with empty search term",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = client.UpdateTerms(ctx, &st.UpdateSearchTermsRequest{Term: term})
	if err != nil {
		log.Printf("error searching for %s, err: %v", term, err)
		c.JSON(400, gin.H{
			"error": fmt.Sprintf("error searching for %s, err: %v", term, err),
		})
	}

	c.JSON(http.StatusOK, gin.H{})
	return
}

func getTopFiveSearchResults(c *gin.Context) {
	conn, err := grpc.Dial(searchtermsbackendaddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("could not connect: %v", err)
		c.JSON(404, gin.H{
			"error": fmt.Sprintf("could not connect: %v", err),
		})
	}
	defer conn.Close()

	client := st.NewSearchClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	searchMetrics, err := client.GetSearchMetrics(ctx, &st.SearchTermMetricsRequest{})
	if err != nil {
		log.Printf("error getting metrics, err: %v", err)
		c.JSON(400, gin.H{
			"error": fmt.Sprintf("error getting metrics, err: %v", err),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"results": searchMetrics.Results,
	})
	return
}
