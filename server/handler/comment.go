package handler

import (
	"database/sql"
	"github.com/Ahmad-mufied/iducate-community-service/server/middlewares"
	"github.com/Ahmad-mufied/iducate-community-service/utils"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

func GetUpdatedCommentCountHandler(c echo.Context) error {
	// Get the post ID from the request parameters
	postIDParam := c.Param("post_id")
	postID, err := strconv.Atoi(postIDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid post ID"})
	}

	// Use the request's context
	ctx := c.Request().Context()

	// Fetch the updated comment count for the post
	count, err := entity.Comment.GetCommentCount(ctx, postID)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Post not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Return the updated comment count as a JSON response
	return c.JSON(http.StatusOK, map[string]int{"comment_count": count})
}

func DeleteCommentHandler(c echo.Context) error {
	userID := middlewares.GetUserID(c)

	// Parse comment ID from URL parameter
	commentIDParam := c.Param("id")
	commentID, err := strconv.Atoi(commentIDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid comment ID"})
	}

	// Use the request's context
	ctx := c.Request().Context()

	// Attempt to delete the comment
	err = entity.Comment.DeleteComment(ctx, uint(commentID), userID)
	if err != nil {
		if err.Error() == "comment not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Comment not found"})
		}
		if err.Error() == "unauthorized: you are not the owner of this comment" {
			return c.JSON(http.StatusForbidden, map[string]string{"error": "You are not authorized to delete this comment"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Return success message
	return c.JSON(http.StatusOK, map[string]string{"message": "Comment deleted successfully"})
}

func CreateCommentHandler(c echo.Context) error {

	userID := middlewares.GetUserID(c)

	// Parse post ID from URL parameter
	postIDParam := c.Param("post_id")
	log.Println(postIDParam)
	postID, err := strconv.Atoi(postIDParam)
	log.Println(postID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid post ID"})
	}

	// Parse user ID and content from request body
	type RequestBody struct {
		Content string `json:"content" validate:"required"`
	}
	var body RequestBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// Validate
	err = validate.Struct(body)
	if err != nil {
		// Format the validation errors
		errors := utils.FormatValidationErrors(err)
		return utils.HandleValidationError(c, errors)
	}

	// Use the request's context
	ctx := c.Request().Context()

	// Create the comment
	comment, err := entity.Comment.CreateComment(ctx, uint(postID), userID, body.Content)
	if err != nil {
		if err.Error() == "post not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Post not found"})
		}
		if err.Error() == "user not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Return the created comment as JSON
	return c.JSON(http.StatusCreated, comment)
}
