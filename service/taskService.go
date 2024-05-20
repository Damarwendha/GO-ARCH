package service

import (
	"go-arch/model"
	"go-arch/model/dto"
	"go-arch/repository"
)

type taskService struct {
	repo repository.TaskRepoI
}

// Create implements TaskServiceI.
func (t *taskService) Create(payload model.Task) error {

	return t.repo.Create(payload)
}

// FindAll implements TaskServiceI.
func (t *taskService) FindAll(page int, size int) ([]model.Task, dto.Paging, error) {

	return t.repo.FindAll(page, size)
}

// FindById implements TaskServiceI.
func (t *taskService) FindById(id string) (model.Task, error) {

	return t.repo.FindById(id)
}

// UpdateById implements TaskServiceI.
func (t *taskService) UpdateById(id string, payload model.Task) error {

	return t.repo.UpdateById(id, payload)
}

type TaskServiceI interface {
	Create(payload model.Task) error
	FindAll(page, size int) ([]model.Task, dto.Paging, error)
	FindById(id string) (model.Task, error)
	UpdateById(id string, payload model.Task) error
}

func NewTaskService(repo repository.TaskRepoI) TaskServiceI {
	return &taskService{
		repo: repo,
	}
}
