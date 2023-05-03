package repository

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

func NewRepository() *Repository {
	return &Repository{}
}
