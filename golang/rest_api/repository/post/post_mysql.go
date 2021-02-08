package post

import (
	"context"
	"crud-api/models"
	"crud-api/repository"
	"crud-api/utils"
	"database/sql"
)

type mysqlPostRepo struct {
	Conn *sql.DB
}

func NewSQLPostRepo(Conn *sql.DB) repository.PostRepo {
	return &mysqlPostRepo{
		Conn: Conn,
	}
}

func (m *mysqlPostRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Post, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.Post, 0)
	for rows.Next() {
		data := new(models.Post)

		err := rows.Scan(
			&data.ID,
			&data.Title,
			&data.Content,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (m *mysqlPostRepo) Fetch(ctx context.Context, num int64) ([]*models.Post, error) {
	query := "SELECT id,title,content FROM posts LIMIT ?"
	return m.fetch(ctx, query, num)
}

func (m *mysqlPostRepo) GetByID(ctx context.Context, id int64) (*models.Post, error) {
	query := "SELECT id,title,content FROM posts WHERE id=?"
	rows, err := m.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}
	payload := &models.Post{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, utils.ErrNotFound
	}

	return payload, err
}

func (m *mysqlPostRepo) Create(ctx context.Context, p *models.Post) (int64, error) {
	query := "INSERT posts SET title=?, content=?"
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}
	res, err := stmt.ExecContext(ctx, p.Title, p.Content)
	defer stmt.Close()
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

func (m *mysqlPostRepo) Update(ctx context.Context, p *models.Post) (*models.Post, error) {
	query := "UPDATE posts SET title=?,content=? WHERE id=?"
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		p.Title,
		p.Content,
		p.ID)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return p, nil
}

func (m *mysqlPostRepo) Delete(ctx context.Context, id int64) (bool, error) {
	query := "DELETE FROM posts WHERE id=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (m *mysqlPostRepo) Init(ctx context.Context) (bool, error) {
	query := "CREATE TABLE IF NOT EXISTS posts (id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, title varchar(100) not null, content varchar(255) not null);"
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}
	_, err = stmt.ExecContext(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}
