package service

import (
	"go-arch/model"
	"go-arch/model/dto"
	"go-arch/repository"
)

type authorService struct {
	repo repository.AuthorRepoI
}

// FindByEmail implements AuthorServiceI.
func (a *authorService) FindByEmail(email string) (model.Author, error) {
	return a.repo.FindByEmail(email)
}

type AuthorServiceI interface {
	FindAll(page, size int) ([]model.Author, dto.Paging, error)
	FindById(id string) (model.Author, error)
	FindByEmail(email string) (model.Author, error)
}

// findAll implements AuthorServiceI.
func (a *authorService) FindAll(page, size int) ([]model.Author, dto.Paging, error) {
	return a.repo.FindAll(page, size)
}

// findById implements AuthorServiceI.
func (a *authorService) FindById(id string) (model.Author, error) {
	return a.repo.FindById(id)
}

func NewAuthorService(repo repository.AuthorRepoI) AuthorServiceI {
	return &authorService{repo: repo}
}
