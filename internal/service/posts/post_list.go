package posts

import (
	"context"
	"github.com/rs/zerolog/log"
	"project2/internal/model/posts"
)

func (s *service) GetAllPost(ctx context.Context, pageSize, pageIndex int) (*posts.GetAllPostResponse, error){
	limit := pageSize
	offset := limit * (pageIndex - 1)
	
	response, err := s.postRepo.GetAllPost(ctx, limit, offset)
	if err != nil{
		log.Error().Err(err).Msg("error getting all post")
		return response, err
	}
	return response, nil
}