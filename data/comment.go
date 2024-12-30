package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/xeonx/timeago"
	"time"
)

type Comment struct {
	ID        uint      `json:"id" db:"id"`                 // Primary key
	PostID    uint      `json:"post_id" db:"post_id"`       // Foreign key referencing Post
	UserID    string    `json:"user_id" db:"user_id"`       // Foreign key referencing User
	Content   string    `json:"content" db:"content"`       // Comment content
	CreatedAt time.Time `json:"created_at" db:"created_at"` // Timestamp for record creation
}

type CommentResponse struct {
	ID        uint   `json:"id" db:"id"`
	Username  string `json:"username" db:"username"`
	Content   string `json:"content" db:"content"`
	CreatedAt string `json:"created_at" db:"created_at"`
}

func (c *Comment) GetComments(ctx context.Context, postID uint) ([]CommentResponse, error) {
	query := `
		SELECT comments.id, users.username, comments.content, comments.created_at
		FROM comments
		JOIN users ON comments.user_id = users.id
		WHERE comments.post_id = $1
		ORDER BY comments.created_at DESC;
	`

	var comments []CommentResponse
	err := db.SelectContext(ctx, &comments, query, postID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch comments: %w", err)
	}

	return comments, nil
}

func (c *Comment) CreateComment(ctx context.Context, postID uint, userID string, content string) (CommentResponse, error) {
	// Check if the post exists
	checkPostQuery := `SELECT id FROM posts WHERE id = $1;`
	var existingPostID uint
	err := db.GetContext(ctx, &existingPostID, checkPostQuery, postID)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return CommentResponse{}, fmt.Errorf("post not found")
		}
		return CommentResponse{}, fmt.Errorf("failed to validate post existence: %w", err)
	}

	// Check if the user exists
	checkUserQuery := `SELECT id FROM users WHERE id = $1;`
	var existingUserID string
	err = db.GetContext(ctx, &existingUserID, checkUserQuery, userID)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return CommentResponse{}, fmt.Errorf("user not found")
		}
		return CommentResponse{}, fmt.Errorf("failed to validate user existence: %w", err)
	}

	// Insert the comment
	insertQuery := `
        INSERT INTO comments (post_id, user_id, content, created_at)
        VALUES ($1, $2, $3, NOW())
        RETURNING id, content, created_at;
    `
	var comment Comment
	err = db.GetContext(ctx, &comment, insertQuery, postID, userID, content)
	if err != nil {
		return CommentResponse{}, fmt.Errorf("failed to create comment: %w", err)
	}

	// convert the comment to comment repsonse
	timestring := timeago.English.Format(comment.CreatedAt)
	commentResponse := CommentResponse{
		ID:        comment.ID,
		Username:  userID,
		Content:   comment.Content,
		CreatedAt: timestring,
	}

	return commentResponse, nil
}

func (c *Comment) DeleteComment(ctx context.Context, commentID uint, userID string) error {
	// Verify that the comment belongs to the user
	checkQuery := `SELECT user_id FROM comments WHERE id = $1;`
	var commentOwnerID string
	err := db.GetContext(ctx, &commentOwnerID, checkQuery, commentID)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return fmt.Errorf("comment not found")
		}
		return fmt.Errorf("failed to verify comment ownership: %w", err)
	}

	if commentOwnerID != userID {
		return fmt.Errorf("unauthorized: you are not the owner of this comment")
	}

	// Delete the comment
	deleteQuery := `DELETE FROM comments WHERE id = $1;`
	result, err := db.ExecContext(ctx, deleteQuery, commentID)
	if err != nil {
		return fmt.Errorf("failed to delete comment: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("comment not found")
	}

	return nil
}

func (c *Comment) GetCommentCount(ctx context.Context, postID int) (int, error) {
	query := `
        SELECT COUNT(*)
        FROM comments
        WHERE post_id = $1;
    `

	var count int
	err := db.GetContext(ctx, &count, query, postID)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch comment count: %w", err)
	}
	return count, nil
}
