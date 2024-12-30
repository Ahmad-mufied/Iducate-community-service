package server

import (
	"github.com/Ahmad-mufied/iducate-community-service/server/handler"
	"github.com/Ahmad-mufied/iducate-community-service/server/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func Routes(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/posts", handler.GetPaginatedPostsHandler) // Get paginated and sorted list of posts
	e.GET("/posts/:id", handler.GetPostDetailHandler)

	e.POST("/posts", handler.CreatePostHandler, middlewares.CognitoJWTMiddleware())       // Create a new post
	e.DELETE("/posts/:id", handler.DeletePostHandler, middlewares.CognitoJWTMiddleware()) // Delete a post by ID

	// Add CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAuthorization},
	}))
	
	// Commnet
	commentGroup := e.Group("/comments")
	commentGroup.GET("/post/:post_id", handler.GetUpdatedCommentCountHandler) // Get comments for a post
	// Comment a post
	commentGroup.POST("/post/:post_id", handler.CreateCommentHandler, middlewares.CognitoJWTMiddleware()) // Get paginated comments for a post
	// Delete a comment
	e.DELETE("/comments/:id", handler.DeleteCommentHandler, middlewares.CognitoJWTMiddleware()) // Delete a comment by ID

	// Like a post
	// Group by like route
	likesGroup := e.Group("/likes")

	likesGroup.GET("/post/:post_id", handler.GetLikesCountHandler)                                     // Get total likes for a post
	likesGroup.POST("/post/:post_id", handler.LikePostHandler, middlewares.CognitoJWTMiddleware())     // Like a post
	likesGroup.DELETE("/post/:post_id", handler.UnlikePostHandler, middlewares.CognitoJWTMiddleware()) // Unlike a post

}
