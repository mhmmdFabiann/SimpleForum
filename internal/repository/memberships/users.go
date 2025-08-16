package memberships

import (
	"context"
	"database/sql"
	"project2/internal/model/memberships"
	"errors"
)

// File: repository/memberships/user.go

func (r *Repository) GetUser(ctx context.Context, email, username string, userID int64) (*memberships.UserModel, error) {
    querry := `SELECT id, email, password, username, created_at, updated_at, created_by, updated_by FROM users WHERE email = ? AND username = ? OR userID = ?`
    row := r.db.QueryRowContext(ctx, querry, email, username, userID)

    var response memberships.UserModel
    err := row.Scan(&response.ID, &response.Email, &response.Password, &response.Username, &response.CreatedAt, &response.UpdatedAt, &response.CreatedBy, &response.UpdatedBy)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) { // Lebih baik menggunakan errors.Is untuk perbandingan error
            return nil, nil // Tidak ditemukan, bukan error sistem
        }
        return nil, err // Error lain dari database
    }
    return &response, nil
}

func(r *Repository) CreateUser(ctx context.Context, model *memberships.UserModel) error {
	query := `INSERT INTO users (email, password, username, created_at, updated_at, created_by, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_ ,err := r.db.ExecContext(ctx, query, model.Email, model.Password, model.Username, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)
	if err != nil{
		return err
	}

	return nil
}