package service

import "restApi/cmd/pkg/repository"

type Authorization interface {
}

type BlogList interface {
}

type BlogItem interface {
}

type Service struct {
	Authorization
	BlogList
	BlogItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
