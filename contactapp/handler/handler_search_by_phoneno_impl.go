package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *impl) SearchByPhoneNo(c *gin.Context) {
	// Get the query parameters from the request
	phoneNo := c.Query("phoneNo")
	requesterPhoneNo := c.Query("requesterPhoneNo")

	contacts, err := h.authService.SearchUserByPhoneNo(c, phoneNo, requesterPhoneNo)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, contacts)
}
