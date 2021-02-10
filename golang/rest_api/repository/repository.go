package repository

import (
	"context"
	"crud-api/models"
)

type PostRepo interface {
	Fetch(ctx context.Context, num int) ([]*models.Post, error)
	GetByID(ctx context.Context, id int64) (*models.Post, error)
	Create(ctx context.Context, p *models.Post) error
	Update(ctx context.Context, p *models.Post) error
	Delete(ctx context.Context, id int) error
}
