package memberships

import (
	"context"
	"errors"

	//"log"

	"project2/internal/model/memberships"

	//"github.com/golang-jwt/jwt/v5"
	"project2/pkg/jwt"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, req memberships.LoginRequest) (string, error){
	// validasi email ada atau tidak

	user, err := s.membershipRepo.GetUser(ctx, req.Email, "")
	if err != nil{ 
		log.Error().Err(err).Msg("fail to get user")
		return "", err
	}
	if user == nil{
		return "", errors.New("email doesnt exsist")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil{
		return "", errors.New("invalid email orpassword")
	}

	token, err := jwt.CreatToken(user.ID, user.Email, s.cfg.Service.SecretJWT)
	if err != nil{
		return "", err
	}
	return token, nil
}