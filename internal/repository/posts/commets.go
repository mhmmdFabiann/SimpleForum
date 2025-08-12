package posts

import (
	"context"
	"project2/internal/model/posts"
)

func (r *Repository) CreateComment(ctx context.Context,model *posts.CommentModel) error{
	querry := `inster into comments(post_id, user_id, comment_content, created_at, updated_at, created_by, updated_by) values (?, ?, ?, ?, ?, ?, ?)`
	_, err:= r.db.ExecContext(ctx, querry, model.PostID, model.UserID , model.CommentContent, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)
	if err != nil{
		return err
	}
	return nil
}