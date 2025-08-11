package posts

import (
	"context"
	"project2/internal/configs"
	"project2/internal/model/posts"
)

func NewService(cfg *configs.Config, postRepo postRepository) *service {
	return &service{
		cfg: cfg,
		postRepo: postRepo,
	}
}

type postRepository interface{
	CreatePost(ctx context.Context, model *posts.PostModel) error
}

type service struct {
	cfg *configs.Config
	postRepo postRepository
}