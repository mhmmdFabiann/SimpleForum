package post

import (
	"github.com/gin-gonic/gin"

 	"project2/internal/model/posts"
 	"project2/internal/middleware"

	"context"
)

type Handler struct {
	*gin.Engine
	postSvc postService
}

type postService interface{
	// mengambil dari service
	CreatePost(ctx context.Context, userID int64 , req *posts.CreatePostRequest) error
	CreateComment(ctx context.Context, postID, userID int64 , req *posts.CreateCommentRequest) error 
	UpsertUserActivity(ctx context.Context, postID, userID int64, req *posts.UserActivityReequest) error
}

func NewHandler(api *gin.Engine, postSvc postService) *Handler{
	return &Handler{
		Engine: api,
		postSvc: postSvc,
	}
}

func(h *Handler) RegisterRoute(){
	route := h.Group("posts")
	route.Use(middleware.AuthMiddleware())

	route.POST("/create", h.CreatePost)
	route.POST("/comment/:postID", h.CreateComment)
	route.PUT("/activity/:postID", h.UpsertUserActivity)
}