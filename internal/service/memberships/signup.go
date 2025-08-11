package memberships

import (
	"context"
	"errors"
	"project2/internal/model/memberships"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (s *service) SingUp(ctx context.Context, req *memberships.SignUpRequest) error{
	user, err := s.membershipRepo.GetUser(ctx, req.Email, req.Username) // cek apakah email dan username sudah ada di database
	if err != nil{
		return err
	}
	if user != nil{
		return errors.New("user already exists")
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil{
		return err
	}

	now := time.Now()
	model := memberships.UserModel{
		Email:     req.Email,
		Password:  string(pass),
		Username:  req.Username,
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: req.Email,
		UpdatedBy: req.Email,
	}
	return s.membershipRepo.CreateUser(ctx, &model)
}