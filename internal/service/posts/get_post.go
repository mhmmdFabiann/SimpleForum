package posts

import (
	"context"
	"project2/internal/model/posts"

	"github.com/rs/zerolog/log"
)

func (s *service) GetPostByID(ctx context.Context, postID int64)(*posts.GetPostResponse, error){
	postDetail, err := s.postRepo.GetPostByID(ctx, postID)
	if err != nil{
		log.Error().Err(err).Msg("error getting post by id")
		return nil, err
	}

	likeCount,err := s.postRepo.CountLikeByPostID(ctx, postID)
	if err != nil{
		log.Error().Err(err).Msg("error counting like by post id")
		return nil, err
	}

	comments, err := s.postRepo.GetCommentByPostID(ctx, postID)
	if err != nil{
		log.Error().Err(err).Msg("error getting comment by post id")
		return nil, err
	}

	return &posts.GetPostResponse{
		PostDetail: posts.Post{
			ID: postDetail.ID,
			UserID: postDetail.UserID,
			Username: postDetail.Username,
			PostTitle: postDetail.PostTitle,
			PostContent: postDetail.PostContent,
			PostHastags: postDetail.PostHastags,
			IsLiked: postDetail.IsLiked,
		},
		LikeCount: likeCount,
		Comment: comments,
	}, nil
}