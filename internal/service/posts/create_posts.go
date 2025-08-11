package posts

import (
	"context"
	"project2/internal/model/posts"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

func (s *service) CreatePost(ctx context.Context, userID int64 , req *posts.CreatePostRequest) error{
	postHastags := strings.Join(req.PostHastags, ",")
	uid := strconv.FormatInt(userID, 10)

	model := posts.PostModel{
		UserID: userID,
		PostTitle: req.PostTitle,
		PostContent: req.PostContent,
		PostHastags: postHastags,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: uid,
		UpdatedBy: uid,
	}

	err := s.postRepo.CreatePost(ctx, &model)
	if err != nil{
		log.Error().Err(err).Msg("error creating post")
		return err
	}
	return nil
}