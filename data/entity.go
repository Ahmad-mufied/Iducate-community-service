package data

import (
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func New(dbPool *sqlx.DB) *Models {
	db = dbPool

	return &Models{
		Post:    &Post{},
		Comment: &Comment{},
		Like:    &Like{},
	}
}

type Models struct {
	Post    PostInterfaces
	Comment CommentInterfaces
	Like    LikeInterfaces
}
