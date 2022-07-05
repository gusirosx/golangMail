package handlers

import (
	"golangMail/services"
	"golangMail/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddLocation receive the a location and saves the coordinates in redis
func ContactUs(ctx *gin.Context) {
	// crate an anonymous struct for driver data.
	var contact types.Mail

	// decode the json request to driver
	if err := ctx.BindJSON(&contact); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//contact.Validate()

	if err := services.ContactUs(contact); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusCreated, gin.H{"success": "Your message was sent."})
}
