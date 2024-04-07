package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *impl) SearchByPhoneNo(c *gin.Context) {
	// Get the query parameters from the request
	phoneNo := c.Query("phoneNo")
	requesterPhoneNo := c.Query("requesterPhoneNo")

	contacts, err := h.contactAppService.SearchUserByPhoneNo(c, phoneNo, requesterPhoneNo)
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
