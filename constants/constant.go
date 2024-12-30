package constants

import (
	"github.com/Ahmad-mufied/iducate-community-service/utils"
	"net/http"
)

const (
	ResponseStatusFailed  = "Failed"
	ResponseStatusSuccess = "Success"
)

var (
	ErrNotFound            = utils.NewAPIError(http.StatusNotFound, "Resource not found", "")
	ErrBadRequest          = utils.NewAPIError(http.StatusBadRequest, "Invalid request data", "")
	ErrInternalServerError = utils.NewAPIError(http.StatusInternalServerError, "Internal Server Error", "")
	ErrUnauthorized        = utils.NewAPIError(http.StatusUnauthorized, "Unauthorized access", "")
	ErrConflict            = utils.NewAPIError(http.StatusConflict, "Resource already exists", "")
)
