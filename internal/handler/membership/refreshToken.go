package membership

import (
	"project2/internal/model/memberships"
	"net/http"
	"github.com/gin-gonic/gin"
)

func (h *Handler) RefreshToken(c *gin.Context){
	ctx := c.Request.Context()

	var request memberships.RefreshTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetInt64("UserID")
	accesToken, err := h.membershipSvc.ValidateRefreshToken(ctx, userID, &request)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := memberships.RefreshTokenResponse{
		AccesToken: accesToken, 
	}
	c.JSON(http.StatusOK, response)
}