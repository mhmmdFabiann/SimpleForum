package memberships

import (
	"context"
	"errors"
	"time"

	//"log"

	"project2/internal/model/memberships"

	//"github.com/golang-jwt/jwt/v5"
	"project2/pkg/jwt"
	tokenUtil "project2/pkg/token"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, req memberships.LoginRequest) (string, string, error){
	// validasi email ada atau tidak

	user, err := s.membershipRepo.GetUser(ctx, req.Email, "", 0)
	if err != nil{ 
		log.Error().Err(err).Msg("fail to get user")
		return "", "",err
	}
	if user == nil{
		return "", "",errors.New("email doesnt exsist")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil{
		return "", "", errors.New("invalid email or password")
	}

	token, err := jwt.CreatToken(user.ID, user.Email, s.cfg.Service.SecretJWT)
	if err != nil{
		return "", "", err
	}

	// sebelum membuat referesh token, lakukan pengecekan apakah sudah ada atau belum
	// jika belum ada, maka buat referesh token
	// jika sudah ada, maka ambil referesh token yang sudah ada
	existingRefreshToken, err := s.membershipRepo.GetRefreshToken(ctx, user.ID, time.Now())
	if err!= nil{
		log.Error().Err(err).Msg("fail to get latest refresh token")
		return "", "", err
	}

	if existingRefreshToken != nil{
		return token, existingRefreshToken.RefreshToken, nil
	}
	
	refreshToken, err:= tokenUtil.GenerateRefreshToken()
	if refreshToken == ""{
		return token, "", err
	}
	err = s.membershipRepo.InsertRefreshToken(ctx, &memberships.RefreshTokenModel{
		UserID: user.ID,
		RefreshToken: refreshToken,
		ExpiredAt: time.Now().Add(7 *24 * time.Hour),
	})
	if err != nil {
		log.Error().Err(err).Msg("fail to insert refresh token")
		return token, refreshToken, err
	}

	return token, refreshToken, nil
}