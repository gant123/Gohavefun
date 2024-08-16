package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.ganthome.cloud/hotel/server"
)

func createHotel(c *gin.Context) {
	var hotel server.Hotel
	if err := c.ShouldBindJSON(&hotel); err == nil {
		// add HATEOS links
		hotel.generateHateosLinks(c.Request.URL.String())
		// add hotel to repository
		server.repository[hotel.Id] = &hotel
		//return OK
		c.JSON(http.StatusAccepted, gin.H{"status": "created"})
	} else {
		// some params not correct
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (h *server.Hotel) generateHateosLinks(url string) {
	// Book url
	postLink := server.Link{
		Href: url + "book",
		Rel:  "book",
		Type: "POST",
	}
	h.Links = append(h.Links, postLink)
}
