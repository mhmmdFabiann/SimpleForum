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
	CreateComment(ctx context.Context,model *posts.CommentModel) error

	GetUserActivity(ctx context.Context, model posts.UserActivityModel) (*posts.UserActivityModel, error)
	CreateUserActivity(ctx context.Context, model *posts.UserActivityModel) error
	UpdateUserActivity(ctx context.Context, model *posts.UserActivityModel) error

	GetAllPost(ctx context.Context, limit, offset int) (*posts.GetAllPostResponse, error)
	GetPostByID(ctx context.Context, id int64) (*posts.Post, error)
	CountLikeByPostID(ctx context.Context, postID int64) (int, error)
	GetCommentByPostID(ctx context.Context, postID int64) ([]*posts.Comment, error)
}

type service struct {
	cfg *configs.Config
	postRepo postRepository
}