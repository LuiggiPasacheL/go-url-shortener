package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/LuiggiPasacheL/go-url-shortener/internal/models"
)

type UrlRepositorySqlite struct {
	db *sql.DB
}

func NewUrlRepositorySqlite(file string) (*UrlRepositorySqlite, error){
	db, err := sql.Open("sqlite3", "file:"+file+".sqlite?cache=shared")
	if err != nil {
		return nil, err
	}

	repository := &UrlRepositorySqlite{
		db: db,
	}

	_, err = db.Exec("IF NOT EXISTS CREATE TABLE URL(id INTEGER PRIMARY KEY AUTOINCREMENT, url VARCHAR(250), shortUrl VARCHAR(250))")
	if err != nil {
		return nil, err
	}

	return repository, nil
}

func (r *UrlRepositorySqlite) Create(ctx context.Context, url *models.Url) (*models.Url, error) {
	return nil, errors.New("Not implemented")
}

func (r *UrlRepositorySqlite) GetById(ctx context.Context, id int) (*models.Url, error) {
	return nil, errors.New("Not implemented")
}

func (r *UrlRepositorySqlite) GetByShortUrl(ctx context.Context, shortUrl string) (*models.Url, error) {
	return nil, errors.New("Not implemented")
}

func (r *UrlRepositorySqlite) GetAll(ctx context.Context) ([]*models.Url, error) {
	return nil, errors.New("Not implemented")
}


func (r *UrlRepositorySqlite) GetByUrl(ctx context.Context, url string) (*models.Url, error) {
	return nil, errors.New("Not implemented")
}

func (r *UrlRepositorySqlite) Close() error {
	if r.db != nil {
		return r.db.Close()
	}
	return nil
}
