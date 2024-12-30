package data

import (
	"context"
	"fmt"
	"github.com/Ahmad-mufied/iducate-community-service/utils"
	"github.com/xeonx/timeago"
	"time"
)

type Post struct {
	ID        uint      `json:"id" db:"id"`                 // Primary key
	UserID    string    `json:"user_id" db:"user_id"`       // Foreign key referencing User
	Title     string    `json:"title" db:"title"`           // Post title
	Content   string    `json:"content" db:"content"`       // Post content
	Views     int       `json:"views" db:"views"`           // Number of views
	CreatedAt time.Time `json:"created_at" db:"created_at"` // Timestamp for record creation
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"` // Timestamp for last update
}

type PostResponse struct {
	ID           uint   `db:"id" json:"id"`
	Title        string `db:"title" json:"title"`
	Content      string `db:"content" json:"content"`
	Views        int    `db:"views" json:"views"`
	Author       string `db:"author" json:"author"`
	LikeCount    int    `db:"like_count" json:"like_count"`
	CommentCount int    `db:"comment_count" json:"comment_count"`
	CreatedAt    string `db:"created_at" json:"created_at"`
}

type PostAndCommentResponse struct {
	ID           uint               `db:"id" json:"id"`
	Title        string             `db:"title" json:"title"`
	Content      string             `db:"content" json:"content"`
	Views        int                `db:"views" json:"views"`
	Author       string             `db:"author" json:"author"`
	LikeCount    int                `db:"like_count" json:"like_count"`
	CommentCount int                `db:"comment_count" json:"comment_count"`
	CreatedAt    string             `db:"created_at" json:"created_at"`
	Comments     []*CommentResponse `json:"comments"`
}

type CreatePostRequest struct {
	UserID  string `json:"user_id" validate:"required"`
	Title   string `json:"title" validate:"required,max=255"`
	Content string `json:"content" validate:"required"`
}

func (p *Post) CreatePost(ctx context.Context, req *CreatePostRequest) (PostResponse, error) {
	query := `
        INSERT INTO posts (user_id, title, content, created_at, updated_at)
        VALUES ($1, $2, $3, NOW(), NOW())
        RETURNING id, user_id, title, content, views, created_at, updated_at;
    `

	var post Post
	err := db.GetContext(ctx, &post, query, req.UserID, req.Title, req.Content)
	if err != nil {
		return PostResponse{}, fmt.Errorf("failed to create post: %w", err)
	}

	// convert the timestamp to a string
	timestring := timeago.English.Format(post.CreatedAt)

	// return the created post with post response
	var postResponse PostResponse
	postResponse.ID = post.ID
	postResponse.Title = post.Title
	postResponse.Content = post.Content
	postResponse.Views = post.Views
	postResponse.Author = post.UserID
	postResponse.CreatedAt = timestring

	return postResponse, nil
}

func (p *Post) GetPaginatedPosts(query PaginatedFeedQuery) ([]PostResponse, error) {
	// Dynamically construct the ORDER BY clause
	orderBy := ""
	switch query.SortType {
	case "trend":
		orderBy = fmt.Sprintf("COUNT(DISTINCT likes.id) %s", query.Sort)
	case "latest":
		orderBy = fmt.Sprintf("MAX(posts.created_at) %s", query.Sort)
	default:
		// This should never happen because `query.SortType` is already validated
		return nil, fmt.Errorf("unexpected sortType: %s", query.SortType)
	}

	// Construct the SQL query
	sqlQuery := fmt.Sprintf(`
        SELECT 
            posts.id,
            posts.title,
            posts.content,
            posts.views,
            users.username AS author,
            COUNT(DISTINCT likes.id) AS like_count,
            COUNT(DISTINCT comments.id) AS comment_count,
            posts.created_at
        FROM posts
        LEFT JOIN likes ON likes.post_id = posts.id
        LEFT JOIN comments ON comments.post_id = posts.id
        JOIN users ON posts.user_id = users.id
        GROUP BY posts.id, users.username, posts.created_at
        ORDER BY %s
        LIMIT $1 OFFSET $2;
    `, orderBy)

	// Execute the query
	var posts []PostResponse
	err := db.Select(&posts, sqlQuery, query.Limit, query.Offset)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch posts: %w", err)
	}

	// Convert the timestamp to a string
	for i := range posts {
		timestamp := posts[i].CreatedAt
		timestring, _ := utils.ParsePostgresTimestamp(timestamp)
		posts[i].CreatedAt = timeago.English.Format(timestring)
	}

	return posts, nil
}

func (p *Post) GetPostDetailWithComments(ctx context.Context, postID uint) (*PostResponse, []*CommentResponse, error) {
	query1 := `
        SELECT
    posts.id,
    posts.title,
    posts.content,
    posts.views,
    users.username AS author,
    COUNT(DISTINCT likes.id) AS like_count,
    COUNT(DISTINCT comments.id) AS comment_count,
    posts.created_at
FROM posts
         LEFT JOIN likes ON likes.post_id = posts.id
         LEFT JOIN comments ON comments.post_id = posts.id
         JOIN users ON posts.user_id = users.id
WHERE posts.id = $1
GROUP BY posts.id, users.username, posts.created_at;
    `

	postDetail := new(PostResponse)
	err := db.GetContext(ctx, postDetail, query1, postID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch post details: %w", err)
	}

	// Convert the timestamp to a string
	timestamp := postDetail.CreatedAt
	timestring, _ := utils.ParsePostgresTimestamp(timestamp)
	postDetail.CreatedAt = timeago.English.Format(timestring)

	query2 := `
SELECT comments.id, users.username, comments.content, comments.created_at
FROM comments
    JOIN posts ON comments.post_id = posts.id
    JOIN users ON comments.user_id = users.id
         WHERE posts.id = $1;`

	var comments []*CommentResponse
	err = db.SelectContext(ctx, &comments, query2, postID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch comments: %w", err)
	}

	// Convert the timestamp to a string
	for i := range comments {
		timestamp := comments[i].CreatedAt
		timestring, _ := utils.ParsePostgresTimestamp(timestamp)
		comments[i].CreatedAt = timeago.English.Format(timestring)
	}

	return postDetail, comments, nil
}

func (p *Post) IncrementPostViews(ctx context.Context, postID uint) error {
	query := `UPDATE posts SET views = views + 1 WHERE id = $1;`

	_, err := db.ExecContext(ctx, query, postID)
	if err != nil {
		return fmt.Errorf("failed to increment views for post ID %d: %w", postID, err)
	}

	return nil
}

func (p *Post) CheckPostByID(ctx context.Context, postID uint) (bool, error) {
	query := `SELECT 1 FROM posts WHERE id = $1;`

	var exists bool
	err := db.GetContext(ctx, &exists, query, postID)
	if err != nil {
		return false, fmt.Errorf("failed to check post existence: %w", err)
	}

	return exists, nil
}

func (p *Post) DeletePost(ctx context.Context, postID uint) error {
	query := `DELETE FROM posts WHERE id = $1;`

	result, err := db.ExecContext(ctx, query, postID)
	if err != nil {
		return fmt.Errorf("failed to delete post: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no post found with the given ID")
	}

	return nil
}
