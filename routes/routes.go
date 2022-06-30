package routes

import (
	"golangMail/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	port = ":8080"
)

func RoutesSetup() {
	// Set Gin to production mode
	//gin.SetMode(gin.ReleaseMode)

	// Set up a http server
	router := gin.Default()

	// Initialize routes
	initializeRoutes(router)

	/*
	   Process the templates at the start so that they don't have to be
	   loaded from the disk again. This makes serving HTML pages very fast.
	*/
	//router.LoadHTMLGlob("templates/*")

	// Run the http server
	if err := router.Run(port); err != nil {
		log.Fatalln("Could not run server: ", err.Error())
	} else {
		log.Println("Starting server on port: ", port)
	}
}

func initializeRoutes(router *gin.Engine) {
	// Handle the index route
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"success": "Up and running..."})
	})
	// Handle the no route case
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Page not found"})
	})

	router.GET("/send", gin.WrapF(service.Index))
	router.POST("/", gin.WrapF(service.Index))

}
