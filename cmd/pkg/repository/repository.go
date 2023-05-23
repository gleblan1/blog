package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type BlogList interface {
}

type BlogItem interface {
}

type Repository struct {
	Authorization
	BlogList
	BlogItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
