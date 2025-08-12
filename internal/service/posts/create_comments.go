package posts

import (
	"context"
	"project2/internal/model/posts"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
)

func (s *service) CreateComment(ctx context.Context, postID, userID int64 , req *posts.CreateCommentRequest) error {
	now := time.Now()
	uid := strconv.FormatInt(userID, 10)

	model := posts.CommentModel{
		PostID: postID,
		UserID: userID,
		CommentContent: req.CommentContent,
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: uid,
		UpdatedBy: uid,
	}

	err := s.postRepo.CreateComment(ctx, &model)
	if err != nil{
		log.Error().Err(err).Msg("error creating comment")
		return err
	}
	return nil
}