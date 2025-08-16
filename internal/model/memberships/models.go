package memberships

import "time"

type (
	SignUpRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Username string `json:"username"`
	}

	LoginRequest struct{
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	RefreshTokenRequest struct{
		Token string `json:"token"`
	}
)

type(
	LoginResponse struct{
		AccesToken string `json:"accesToken"`
		RefreshToken string `json:"refreshToken"`
	}

	RefreshTokenResponse struct{
		AccesToken string `json:"accesToken"`
	}
)

type (
	UserModel struct {
		ID        int64  `db:"id"`
		Email     string `db:"email"`
		Password  string `db:"password"`
		Username  string `db:"username"`
		CreatedAt time.Time `db:"created_at"`
		CreatedBy string `db:"created_by"`
		UpdatedAt time.Time `db:"updated_at"`
		UpdatedBy string `db:"updated_by"`
	}

	RefreshTokenModel struct{
		ID int64 `db:"id"`
		UserID int64 `db:"user_id"`
		RefreshToken string `db:"refresh_token"`
		ExpiredAt time.Time `db:"expired_at"`
		CreatedAt time.Time `db:"created_at"`
		CreatedBy string `db:"created_by"`
		UpdatedAt time.Time `db:"updated_at"`
		UpdatedBy string `db:"updated_by"`
	}
)