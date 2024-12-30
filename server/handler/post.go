package handler

import (
	"github.com/Ahmad-mufied/iducate-community-service/data"
	"github.com/Ahmad-mufied/iducate-community-service/server/middlewares"
	"github.com/Ahmad-mufied/iducate-community-service/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func GetPaginatedPostsHandler(c echo.Context) error {
	// Parse query parameters
	var query data.PaginatedFeedQuery
	if err := query.Parse(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid query parameters"})
	}

	posts, err := entity.Post.GetPaginatedPosts(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Return the response
	return c.JSON(http.StatusOK, posts)
}

func GetPostDetailHandler(c echo.Context) error {
	// Get post ID from URL parameter
	postIDParam := c.Param("id")
	postID, err := strconv.Atoi(postIDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid post ID"})
	}

	// Use the request's context
	ctx := c.Request().Context()

	// Check if the post exists
	exists, err := entity.Post.CheckPostByID(ctx, uint(postID))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Post not found"})
	}

	if !exists {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Post not found"})
	}

	// Fetch post details with comments
	post, comments, err := entity.Post.GetPostDetailWithComments(ctx, uint(postID))
	if err != nil {
		if err.Error() == "post not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Post not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Increment views count
	err = entity.Post.IncrementPostViews(ctx, uint(postID))
	if err != nil {
		if err.Error() == "post not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Post not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	if post == nil {
		post = &data.PostResponse{}
	}

	if comments == nil {
		comments = []*data.CommentResponse{}
	}

	// Define the response structure
	postComment := &data.PostAndCommentResponse{
		ID:           post.ID,
		Title:        post.Title,
		Content:      post.Content,
		Views:        post.Views,
		Author:       post.Author,
		LikeCount:    post.LikeCount,
		CommentCount: post.CommentCount,
		CreatedAt:    post.CreatedAt,
		Comments:     comments,
	}

	// Combine the post and comments into a single response
	//response := struct {
	//	data.PostResponse `json:"post"`
	//	Comments          []*data.CommentResponse `json:"comments"`
	//}{
	//	PostResponse: *post,
	//	Comments:     comments,
	//}

	// Return response
	return c.JSON(http.StatusOK, postComment)
}

func CreatePostHandler(c echo.Context) error {

	userID := middlewares.GetUserID(c)

	var req = new(data.CreatePostRequest)
	req.UserID = userID

	// Bind and validate the request body
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}
	// Validate
	err := validate.Struct(req)
	if err != nil {
		// Format the validation errors
		errors := utils.FormatValidationErrors(err)
		return utils.HandleValidationError(c, errors)
	}

	// Use the request's context
	ctx := c.Request().Context()

	// Create the post
	post, err := entity.Post.CreatePost(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Return the created post as JSON
	return c.JSON(http.StatusCreated, post)
}

func DeletePostHandler(c echo.Context) error {
	// Get post ID from URL parameter
	postIDParam := c.Param("id")
	postID, err := strconv.Atoi(postIDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid post ID"})
	}

	// Use the request's context
	ctx := c.Request().Context()

	// Check if the post exists
	exists, err := entity.Post.CheckPostByID(ctx, uint(postID))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Post not found"})
	}

	if !exists {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Post not found"})
	}

	// Delete the post
	err = entity.Post.DeletePost(ctx, uint(postID))
	if err != nil {
		if err.Error() == "no post found with the given ID" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Post not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Return success message
	return c.JSON(http.StatusOK, map[string]string{"message": "Post deleted successfully"})
}
