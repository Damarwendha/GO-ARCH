package model

type Task struct {
	Id        string `json:"id"`
	Title     string `json:"title" binding:"required"`
	Content   string `json:"content" binding:"required"`
	AuthorId  string `json:"author_id" binding:"required"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
