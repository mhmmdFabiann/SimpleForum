package membership

import "project2/internal/model/memberships"
import "net/http"
import "github.com/gin-gonic/gin"
//import "project2/internal/service/memberships"

func (h *Handler) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var request memberships.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accesToken ,err := h.membershipSvc.Login(ctx, request)
	if  err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := memberships.LoginResponse{
		AccesToken: accesToken,
	}
	c.JSON(http.StatusOK, response)
}