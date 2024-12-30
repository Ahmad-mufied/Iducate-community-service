package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
)

// ExtractIDToken extracts the `id_token` from the request headers.
func ExtractIDToken(r *http.Request) (string, error) {
	idToken := r.Header.Get("id_token")
	if idToken == "" {
		return "", errors.New("missing id_token in headers")
	}
	return idToken, nil
}

// DecodeIDToken decodes the JWT `id_token` without verifying the signature.
func DecodeIDToken(idToken string) (map[string]interface{}, error) {
	// Parse the token without verifying the signature
	token, _, err := jwt.NewParser().ParseUnverified(idToken, jwt.MapClaims{})
	if err != nil {
		return nil, errors.New("invalid id_token")
	}

	// Extract claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("unable to parse claims")
	}

	// Convert claims to a map
	decodedClaims := make(map[string]interface{})
	for key, value := range claims {
		decodedClaims[key] = value
	}

	return decodedClaims, nil
}

// GetUserID extracts the `sub` claim (user ID) from the decoded token claims.
func GetUserID(idToken string) (string, error) {
	claims, err := DecodeIDToken(idToken)
	if err != nil {
		return "", err
	}

	sub, ok := claims["sub"].(string)
	if !ok || strings.TrimSpace(sub) == "" {
		return "", errors.New("sub claim is missing or invalid")
	}

	return sub, nil
}
