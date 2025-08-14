package posts

import (
	"context"
	"project2/internal/model/posts"
	"errors"
	"database/sql"
)

func (r *Repository) GetUserActivity(ctx context.Context, model posts.UserActivityModel) (*posts.UserActivityModel, error){
	querry := `select id, post_id, user_id is_liked, created_at, updated_at, created_by, updated_by from user_activities where post_id = ? and user_id = ?`

	var response posts.UserActivityModel
	row := r.db.QueryRowContext(ctx, querry, model.PostID, model.UserID)

	err:= row.Scan(&response.ID, &response.PostID, &response.IsLiked, &response.CreatedAt, &response.UpdatedAt, &response.CreatedBy, &response.UpdatedBy)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) { // Lebih baik menggunakan errors.Is untuk perbandingan error
            return nil, nil // Tidak ditemukan, bukan error sistem
        }
        return nil, err // Error lain dari database
    }
    return &response, nil
}

func(r *Repository) CreateUserActivity(ctx context.Context, model *posts.UserActivityModel) error{
	querry := `INSERT INTO user_activities (post_id, user_id, is_liked, created_at, updated_at, created_by, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?)`

	_ ,err := r.db.ExecContext(ctx, querry, model.PostID, model.UserID, model.IsLiked, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)
	if err != nil{
		return err	
	}

	return nil
}

func(r *Repository) UpdateUserActivity(ctx context.Context, model *posts.UserActivityModel) error{
	querry := `update user_activities set is_liked = ? where post_id = ? and user_id = ?`

	_ ,err := r.db.ExecContext(ctx, querry, model.IsLiked, model.PostID, model.UserID)
	if err != nil{
		return err
	}

	return nil
}

func (r *Repository) CountLikeByPostID(ctx context.Context, postID int64) (int, error){
	querry := `select count(id) from user_activities where post_id = ? and is_liked = true`

	var response int
	row := r.db.QueryRowContext(ctx, querry, postID)

	err:= row.Scan(&response)
    if err != nil {
		return response, err
    }
    return response, nil
}

func (r *Repository) GetCommentByPostID(ctx context.Context, postID int64) ([]*posts.Comment, error){
	querry := `select c.id, c.user_id, c.comment_content, u.username
	from comments c join users u on c.user_id = u.id
	where c.post_id = ?`

	rows, err := r.db.QueryContext(ctx, querry, postID)
	if err != nil{
		return nil, err
	}
	defer rows.Close()
	response := make([]*posts.Comment, 0)

	for rows.Next(){
		var(
			comment posts.Comment
			username string
		)
		err := rows.Scan(&comment.ID, &comment.UserID, &comment.CommentContent, &username)
		if err != nil{
			return nil, err
		}
		response = append(response, &posts.Comment{
			ID: comment.ID,
			UserID: comment.UserID,
			CommentContent: comment.CommentContent,
			Username: username,
		})
	}
	return response, nil
}