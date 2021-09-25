package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.RedirectFixedPath = true
	router.GET("/search-term/:term", searchTerms)
	router.PUT("/update-search-terms/:term", updateSearchTerms)
	router.GET("/search-term-metrics", getTopFiveSearchResults)
	router.Run()
}
