package middlewares

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"strings"
)

// CognitoJWTMiddleware extracts and parses the Cognito ID token from headers
func CognitoJWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Extract ID token from headers
			idToken := c.Request().Header.Get("id_token")
			if idToken == "" {
				return echo.NewHTTPError(401, "Missing id_token in headers")
			}

			// Decode token without signature verification
			claims, err := decodeIDToken(idToken)
			if err != nil {
				return echo.NewHTTPError(401, err.Error())
			}

			// Set token claims in context for later use
			c.Set("token_claims", claims)
			c.Set("user_id", claims["sub"])
			c.Set("name", claims["name"])

			return next(c)
		}
	}
}

// decodeIDToken decodes the JWT token without signature verification
func decodeIDToken(tokenString string) (jwt.MapClaims, error) {
	// Remove "Bearer " prefix if present
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// Parse token without signature verification
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return nil, errors.New("invalid id_token")
	}

	// Type assert to get claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("unable to extract claims")
	}

	// Optional: Check token expiration manually
	//if exp, ok := claims["exp"].(float64); ok {
	//	if jwt.NewNumericDate(time.Now()).Unix() > int64(exp) {
	//		return nil, errors.New("id_token has expired")
	//	}
	//}

	return claims, nil
}

// Helper function to get user ID from context
func GetUserID(c echo.Context) string {
	userID, ok := c.Get("user_id").(string)
	if !ok {
		return ""
	}
	return userID
}

// Helper function to get username from context
func GetUsername(c echo.Context) string {
	username, ok := c.Get("name").(string)
	if !ok {
		return ""
	}
	return username
}
