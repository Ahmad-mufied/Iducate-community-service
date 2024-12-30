package data

import (
	"context"
	"fmt"
)

type Like struct{}

func (l *Like) AddLike(ctx context.Context, userID string, postID int) error {
	query := `
        INSERT INTO likes (user_id, post_id, created_at)
        VALUES ($1, $2, NOW())
        ON CONFLICT DO NOTHING; -- Avoid duplicate likes
    `
	_, err := db.ExecContext(ctx, query, userID, postID)
	if err != nil {
		return fmt.Errorf("failed to add like: %w", err)
	}
	return nil
}

func (l *Like) RemoveLike(ctx context.Context, userID string, postID int) error {
	query := `DELETE FROM likes WHERE user_id = $1 AND post_id = $2;`
	_, err := db.ExecContext(ctx, query, userID, postID)
	if err != nil {
		return fmt.Errorf("failed to remove like: %w", err)
	}
	return nil
}

func (l *Like) CountLikes(ctx context.Context, postID int) (int, error) {
	query := `
        SELECT COUNT(*) 
        FROM likes
        WHERE post_id = $1;
    `
	var likeCount int
	err := db.GetContext(ctx, &likeCount, query, postID)
	if err != nil {
		return 0, fmt.Errorf("failed to count likes: %w", err)
	}
	return likeCount, nil
}
