package memberships

import (
	"context"
	"database/sql"
	"time"

	//"database/sql"
	"project2/internal/model/memberships"
	//"errors"
)

func(r *Repository) InsertRefreshToken(ctx context.Context, model *memberships.RefreshTokenModel) error{
    querry := `INSERT INTO refresh_tokens (user_id, refresh_token,expire_at, created_at, updated_at, created_by, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?)`

	_ ,err := r.db.ExecContext(ctx, querry, model.UserID, model.RefreshToken, model.ExpiredAt, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)
	if err != nil{
		return err
	}

	return nil
}

func(r *Repository) GetRefreshToken(ctx context.Context, UserID int64, now time.Time) (*memberships.RefreshTokenModel, error){
	querry:= `SELECT id, user_id, refresh_token, expire_at, created_at, updated_at, created_by, updated_by FROM refresh_tokens WHERE user_id = ? AND expire_at >= ?`

	var response memberships.RefreshTokenModel
	row := r.db.QueryRowContext(ctx, querry, UserID, now)
	err := row.Scan(&response.ID, &response.UserID, &response.RefreshToken, &response.ExpiredAt, &response.CreatedAt, &response.UpdatedAt, &response.CreatedBy, &response.UpdatedBy)
	if err != nil{
		if err == sql.ErrNoRows{
			return nil, nil
		}
		return nil, err
	}
	return &response, nil
}