package handler

import (
	"github.com/Ahmad-mufied/iducate-community-service/data"
	"github.com/go-playground/validator/v10"
)

var entity *data.Models
var validate *validator.Validate

func InitHandler(m *data.Models, v *validator.Validate) {
	entity = m
	validate = v
}
