package repository

import (
	"database/sql"
	"go-arch/model"
	"go-arch/model/dto"
)

type taskRepo struct {
	db *sql.DB
}

// Create implements TaskRepoI.
func (t *taskRepo) Create(payload model.Task) error {

	stmt := "INSERT INTO trx_tasks (title, content, author_id) VALUES ($1, $2, $3);"

	_, err := t.db.Exec(stmt, payload.Title, payload.Content, payload.AuthorId)
	if err != nil {
		return err
	}

	return nil
}

// FindAll implements TaskRepoI.
func (t *taskRepo) FindAll(page int, size int) ([]model.Task, dto.Paging, error) {
	var listData []model.Task

	offset := (size * page) - size
	rows, err := t.db.Query("SELECT * FROM trx_tasks LIMIT $1 OFFSET $2;", size, offset)

	if err != nil {
		return nil, dto.Paging{}, err
	}

	for rows.Next() {
		var task model.Task

		err := rows.Scan(&task.Id, &task.Title, &task.Content, &task.AuthorId, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, dto.Paging{}, err
		}

		listData = append(listData, task)
	}

	if err = rows.Err(); err != nil {
		return nil, dto.Paging{}, err
	}

	var totalRows int
	err = t.db.QueryRow("SELECT COUNT(*) FROM trx_tasks;").Scan(&totalRows)

	if err != nil {
		return nil, dto.Paging{}, err
	}

	paging := dto.Paging{
		Page:       page,
		Size:       size,
		TotalRows:  totalRows,
		TotalPages: totalRows / size,
	}

	return listData, paging, nil
}

// FindById implements TaskRepoI.
func (t *taskRepo) FindById(id string) (model.Task, error) {
	var task model.Task

	stmt := "SELECT * FROM trx_tasks WHERE id = $1"

	err := t.db.QueryRow(stmt, id).Scan(&task.Id, &task.Title, &task.Content, &task.AuthorId, &task.CreatedAt, &task.UpdatedAt)

	if err != nil {
		return model.Task{}, err
	}

	return task, nil

}

// UpdateById implements TaskRepoI.
func (t *taskRepo) UpdateById(id string, payload model.Task) error {

	stmt := "UPDATE trx_tasks SET title = $1, content = $2, author_id = $3 WHERE id = $4;"

	_, err := t.db.Exec(stmt, payload.Title, payload.Content, payload.AuthorId, id)

	if err != nil {
		return err
	}

	return nil
}

type TaskRepoI interface {
	FindAll(page, size int) ([]model.Task, dto.Paging, error)
	FindById(id string) (model.Task, error)
	Create(payload model.Task) error
	UpdateById(id string, payload model.Task) error
}

func NewTaskRepo(db *sql.DB) TaskRepoI {
	return &taskRepo{db: db}
}
