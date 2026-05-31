package tasks_transport_http

import (
	"time"

	"github.com/Sergey-tech9087/petProjectToDoList/internal/core/domain"
)

type TasksDTOResponse struct {
	ID           int        `json:"id" example:"15"`
	Version      int        `json:"version" example:"3"`
	Title        string     `json:"title" example:"Pet project"`
	Description  *string    `json:"description" example:"Rework pet project add feature"`
	Completed    bool       `json:"completed" example:"false"`
	CreatedAt    time.Time  `json:"created_at" example:"2026-05-31T10:30:00Z"`
	CompletedAt  *time.Time `json:"completed_at" example:"null"`
	AuthorUserID int        `json:"author_user_id" example:"5"`
}

func taskDTOFromDomain(task domain.Task) TasksDTOResponse {
	return TasksDTOResponse{
		ID:           task.ID,
		Version:      task.Version,
		Title:        task.Title,
		Description:  task.Description,
		Completed:    task.Completed,
		CreatedAt:    task.CreatedAt,
		CompletedAt:  task.CompletedAt,
		AuthorUserID: task.AuthorUserID,
	}
}

func tasksDTOsFromDomain(tasks []domain.Task) []TasksDTOResponse {
	dtos := make([]TasksDTOResponse, len(tasks))

	for i, task := range tasks {
		dtos[i] = taskDTOFromDomain(task)
	}

	return dtos
}
