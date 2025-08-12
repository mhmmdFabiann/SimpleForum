package post

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project2/internal/model/posts"
	"strconv"
)

func (h *Handler) UpsertUserActivity(c *gin.Context){
	ctx := c.Request.Context()

	var request posts.UserActivityReequest
	if err := c.ShouldBindJSON(&request); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uid := c.GetInt64("UserID")
	postIDStr := c.Param("postID") // "postID" harus cocok dengan nama di route, misal: /posts/:postID/comments
	pid, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Post ID"})
		return
	}	
	
	err = h.postSvc.UpsertUserActivity(ctx, pid, uid, &request)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}