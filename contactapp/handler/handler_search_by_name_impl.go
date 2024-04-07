package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *impl) SearchByNameHandler(c *gin.Context) {
	// Get the query parameters from the request
	name := c.Query("name")
	requesterPhoneNo := c.Query("requesterPhoneNo")

	// Call the service method to perform the search
	contacts, err := h.contactAppService.SearchByName(c, name, requesterPhoneNo)
	if err != nil {
		if errors.Is(err, errors.New("Unauthorized: User Not Found")) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, contacts)
}
