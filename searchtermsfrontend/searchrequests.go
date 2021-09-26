package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

const (
	searchtermsbackendaddress = "backend:50080"
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

	client := NewSearchClient(conn)

	term := c.Param("term")

	request := SearchTermRequest{
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

	client := NewSearchClient(conn)

	term := c.Param("term")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = client.UpdateTerms(ctx, &UpdateSearchTermsRequest{Term: term})
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

	client := NewSearchClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	searchMetrics, err := client.GetSearchMetrics(ctx, &SearchTermMetricsRequest{})
	if err != nil {
		log.Printf("error getting metrics, err: %v", err)
		c.JSON(400, gin.H{
			"error": fmt.Sprintf("error getting metrics, err: %v", err),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"results": searchMetrics.Results,
	})
}
