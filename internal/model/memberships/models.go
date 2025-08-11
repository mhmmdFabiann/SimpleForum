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
)

type(
	LoginResponse struct{
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
)