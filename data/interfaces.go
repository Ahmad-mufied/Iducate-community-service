package data

import (
	"context"
)

type PostInterfaces interface {
	CreatePost(ctx context.Context, req *CreatePostRequest) (PostResponse, error)
	GetPaginatedPosts(query PaginatedFeedQuery) ([]PostResponse, error)
	GetPostDetailWithComments(ctx context.Context, postID uint) (*PostResponse, []*CommentResponse, error)
	CheckPostByID(ctx context.Context, postID uint) (bool, error)
	IncrementPostViews(ctx context.Context, postID uint) error
	DeletePost(ctx context.Context, postID uint) error
}

type CommentInterfaces interface {
	GetComments(ctx context.Context, postID uint) ([]CommentResponse, error)
	GetCommentCount(ctx context.Context, postID int) (int, error)
	CreateComment(ctx context.Context, postID uint, userID string, content string) (CommentResponse, error)
	DeleteComment(ctx context.Context, commentID uint, userID string) error
}

type LikeInterfaces interface {
	AddLike(ctx context.Context, userID string, postID int) error
	RemoveLike(ctx context.Context, userID string, postID int) error
	CountLikes(ctx context.Context, postID int) (int, error)
}
