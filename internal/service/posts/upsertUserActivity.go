package posts

import (
	"context"
	"errors"
	"project2/internal/model/posts"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
)

func (s *service) UpsertUserActivity(ctx context.Context, postID, userID int64, req *posts.UserActivityReequest) error{
	model := posts.UserActivityModel{
		PostID: postID,
		UserID: userID,
		IsLiked: req.IsLiked,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: strconv.FormatInt(userID, 10),
		UpdatedBy: strconv.FormatInt(userID, 10),
	}
	userActivity, err := s.postRepo.GetUserActivity(ctx, model)
	if err != nil{
		log.Error().Err(err).Msg("error getting user activity")
		return err
	}
	if userActivity == nil{
		//create user activity
		if !req.IsLiked{
			return errors.New("anda belum menyukai post ini")
		}
		err = s.postRepo.CreateUserActivity(ctx, &model)
	}else{
		err=s.postRepo.UpdateUserActivity(ctx, &model)
	}
	if err!=nil{
		log.Error().Err(err).Msg("error create or updating user activity")
		return err
	}
	return nil
}