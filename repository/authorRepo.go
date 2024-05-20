package repository

import (
	"database/sql"
	"go-arch/model"
	"go-arch/model/dto"
)

type authorRepo struct {
	db *sql.DB
}

// FindByEmail implements AuthorRepoI.
func (a *authorRepo) FindByEmail(email string) (model.Author, error) {
	var author model.Author

	stmt := "SELECT * FROM mst_authors WHERE email = $1"

	err := a.db.QueryRow(stmt, email).Scan(&author.Id, &author.Fullname, &author.Email, &author.Password, &author.Created_at, &author.Updated_at, &author.Role)

	if err != nil {
		return model.Author{}, err
	}

	return author, nil
}

type AuthorRepoI interface {
	FindAll(page, size int) ([]model.Author, dto.Paging, error)
	FindById(id string) (model.Author, error)
	FindByEmail(email string) (model.Author, error)
}

// findAll implements AuthorRepoI.
func (a *authorRepo) FindAll(page, size int) ([]model.Author, dto.Paging, error) {
	var listData []model.Author

	offset := (size * page) - size
	rows, err := a.db.Query("SELECT * FROM mst_authors LIMIT $1 OFFSET $2;", size, offset)

	if err != nil {
		return nil, dto.Paging{}, err
	}

	for rows.Next() {
		var author model.Author

		err := rows.Scan(&author.Id, &author.Fullname, &author.Email, &author.Password, &author.Created_at, &author.Updated_at, &author.Role)
		if err != nil {
			return nil, dto.Paging{}, err
		}

		listData = append(listData, author)
	}

	if err = rows.Err(); err != nil {
		return nil, dto.Paging{}, err
	}

	var totalRows int
	err = a.db.QueryRow("SELECT COUNT(*) FROM mst_authors;").Scan(&totalRows)

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

// findById implements AuthorRepoI.
func (a *authorRepo) FindById(id string) (model.Author, error) {
	var author model.Author

	stmt := "SELECT * FROM mst_authors WHERE id = $1"

	err := a.db.QueryRow(stmt, id).Scan(&author.Id, &author.Fullname, &author.Email, &author.Password, &author.Created_at, &author.Updated_at, &author.Role)

	if err != nil {
		return model.Author{}, err
	}

	return author, nil

}

func NewAuthorRepo(db *sql.DB) AuthorRepoI {
	return &authorRepo{db: db}
}
