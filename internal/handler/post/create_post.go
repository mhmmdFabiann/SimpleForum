package post

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project2/internal/model/posts"
)

func (h *Handler) CreatePost(c *gin.Context){
	ctx := c.Request.Context()

	var request posts.CreatePostRequest
	if err := c.ShouldBindJSON(&request); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	uid := c.GetInt64("UserID")
	err := h.postSvc.CreatePost(ctx, uid, &request)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}