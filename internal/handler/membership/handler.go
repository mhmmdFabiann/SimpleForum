package membership

import (
	"github.com/gin-gonic/gin"
 	"project2/internal/model/memberships"
 	"context"
)

type Handler struct {
	*gin.Engine
	membershipSvc membershipService
}

type membershipService interface{
	SingUp(ctx context.Context, req *memberships.SignUpRequest) error
	Login(ctx context.Context, req memberships.LoginRequest) (string, error)
}

func NewHandler(api *gin.Engine, membershipSvc membershipService) *Handler{
	return &Handler{
		Engine: api,
		membershipSvc: membershipSvc,
	}
}

func(h *Handler) RegisterRoute(){
	route := h.Group("membership")
	route.POST("/signup", h.SignUp)
	route.POST("/login", h.Login)
}