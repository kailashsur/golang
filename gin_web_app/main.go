// main.go

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func showIndexPage(c *gin.Context) {

	// call teh html method of the context to render a template
	c.HTML(
		// set the http status to 200 (ok)
		http.StatusOK,

		//User the index.html template
		"index.html",

		// Pass the data that the page use in this case "title"
		gin.H{
			"title": "Home Page",
		},
	)
}

func main() {

	// default way to create router in gin
	router = gin.Default()

	// load the templates
	router.LoadHTMLGlob("templates/*")

	// define the router handler
	router.GET("/", showIndexPage)

	// Start serving the application
	router.Run()

}
