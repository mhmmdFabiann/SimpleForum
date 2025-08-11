package memberships

import (
	"context"
	"project2/internal/configs"
	"project2/internal/model/memberships"
)

func NewService(cfg *configs.Config, membershipRepo membershipRepository) *service {
	return &service{
		cfg: cfg,
		membershipRepo: membershipRepo,
	}
}

type membershipRepository interface{
	GetUser(ctx context.Context, email, username string) (*memberships.UserModel, error)
	CreateUser(ctx context.Context, model *memberships.UserModel) error
}

type service struct {
	cfg *configs.Config
	membershipRepo membershipRepository
}