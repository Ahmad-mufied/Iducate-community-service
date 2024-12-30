package handler

import (
	"github.com/Ahmad-mufied/iducate-community-service/server/middlewares"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func LikePostHandler(c echo.Context) error {

	// Get the user ID from middleware
	userID := middlewares.GetUserID(c)

	// Get the post ID from the request
	postIDParam := c.Param("post_id") // Assume post ID is passed as a route parameter

	postID, err := strconv.Atoi(postIDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid post ID"})
	}

	// Use the request's context
	ctx := c.Request().Context()

	// Add like
	err = entity.Like.AddLike(ctx, userID, postID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Return success message
	return c.JSON(http.StatusOK, map[string]string{"message": "Post liked successfully"})
}

func UnlikePostHandler(c echo.Context) error {
	// Get the user ID from middleware
	userID := middlewares.GetUserID(c)

	// Get the user ID and post ID from the request
	postIDParam := c.Param("post_id") // Assume post ID is passed as a route parameter

	postID, err := strconv.Atoi(postIDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid post ID"})
	}

	// Use the request's context
	ctx := c.Request().Context()

	// Remove like
	err = entity.Like.RemoveLike(ctx, userID, postID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Return success message
	return c.JSON(http.StatusOK, map[string]string{"message": "Post unliked successfully"})
}

func GetLikesCountHandler(c echo.Context) error {
	// Get the post ID from the request
	postIDParam := c.Param("post_id") // Assume post ID is passed as a route parameter

	postID, err := strconv.Atoi(postIDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid post ID"})
	}

	// Use the request's context
	ctx := c.Request().Context()

	// Count likes
	likeCount, err := entity.Like.CountLikes(ctx, postID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Return like count
	return c.JSON(http.StatusOK, map[string]int{"like_count": likeCount})
}
