package post

import (
	"context"
	"crud-api/models"
	"crud-api/repository"
	"errors"
	"gorm.io/gorm"
)

type PostsRepository struct {
	db *gorm.DB
}

func NewSQLPostRepo(db *gorm.DB) repository.PostRepo {
	return &PostsRepository{
		db: db,
	}
}

func (pr *PostsRepository) Fetch(ctx context.Context, num int) ([]*models.Post, error) {
	var p []*models.Post
	pr.db.WithContext(ctx).Limit(num).Find(&p)
	return p, nil
}

func (pr *PostsRepository) GetByID(ctx context.Context, id int64) (*models.Post, error) {
	var p models.Post
	if err := pr.db.WithContext(ctx).First(&p, id); err != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			return nil, err.Error
		}
		return nil, err.Error
	}
	return &p, nil
}

func (pr *PostsRepository) Create(ctx context.Context, p *models.Post) error {
	return pr.db.WithContext(ctx).Create(p).Error
}

//UpdateOrCreate actually...
func (pr *PostsRepository) Update(ctx context.Context, p *models.Post) error {
	res := pr.db.WithContext(ctx).Model(p).Updates(p)
	if res.RowsAffected == 0 {
		err := pr.db.WithContext(ctx).Create(p)
		if err != nil {
			return err.Error
		}
	}
	return nil
}

func (pr *PostsRepository) Delete(ctx context.Context, id int) error {
	var p models.Post
	return pr.db.WithContext(ctx).Delete(&p, id).Error
}
