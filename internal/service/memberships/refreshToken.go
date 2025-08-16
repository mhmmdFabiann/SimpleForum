package memberships

import (
	"context"
	"github.com/rs/zerolog/log"
	"project2/internal/model/memberships"
	"time"
	"errors"
	"project2/pkg/jwt"
)

//return valuenya adalah string acces token dan error
func (s *service) ValidateRefreshToken(ctx context.Context, userID int64, request *memberships.RefreshTokenRequest) (string, error) {
	exsistingRefreshToken, err := s.membershipRepo.GetRefreshToken(ctx, userID, time.Now())
	if err != nil{
		log.Error().Err(err).Msg("fail to get latest refresh token")
		return "", err
	}

	if exsistingRefreshToken == nil{
		return "", errors.New("refresh token has expired")
	}

	if exsistingRefreshToken.RefreshToken != request.Token{
		return "", errors.New("refresh token Is invalid")
	}

	user, err := s.membershipRepo.GetUser(ctx, "", "", userID)
	if err != nil{ 
		log.Error().Err(err).Msg("fail to get user")
		return "", err
	}
	if user == nil{
		return "", errors.New("user doesnt exsist")
	}

	token, err := jwt.CreatToken(user.ID, user.Email, s.cfg.Service.SecretJWT)
	if err != nil{
		return "", err
	}
	return token, nil
}