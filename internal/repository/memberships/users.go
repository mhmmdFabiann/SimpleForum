package memberships

import (
	"context"
	"database/sql"
	"project2/internal/model/memberships"
	"errors"
)

// File: repository/memberships/user.go

func (r *Repository) GetUser(ctx context.Context, email, username string) (*memberships.UserModel, error) {
    // Mulai query dengan kondisi yang selalu benar
    query := `SELECT id, email, password, username, created_at, updated_at, created_by, updated_by FROM users WHERE 1=1`
    // Siapkan slice untuk menampung argumen secara dinamis
    args := []interface{}{}
    // Tambahkan kondisi ke query HANYA JIKA parameternya diisi
    if email != "" {
        query += " AND email = ?" // Ganti ? dengan $1, $2, dst. jika Anda pakai PostgreSQL
        args = append(args, email)
    }
    if username != "" {
        query += " AND username = ?"
        args = append(args, username)
    }
    // Pastikan hanya satu hasil yang diambil
    query += " LIMIT 1"
    // Jalankan query dengan argumen yang sudah dinamis
    row := r.db.QueryRowContext(ctx, query, args...)

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