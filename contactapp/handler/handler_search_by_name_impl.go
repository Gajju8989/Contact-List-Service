package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *impl) SearchByNameHandler(c *gin.Context) {
	// Get the query parameters from the request
	name := c.Query("name")
	requesterPhoneNo := c.Query("requesterPhoneNo")

	// Call the service method to perform the search
	contacts, err := h.authService.SearchByName(c, name, requesterPhoneNo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, contacts)
}
