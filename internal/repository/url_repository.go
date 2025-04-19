package repository

import (
	"context"

	"github.com/LuiggiPasacheL/go-url-shortener/internal/models"
)

type UrlRepository interface {
    Create(ctx context.Context, url *models.Url) (*models.Url, error)
    GetById(ctx context.Context, id int) (*models.Url, error)
    GetByShortUrl(ctx context.Context, shortUrl string) (*models.Url, error)
    GetAll(ctx context.Context) ([]*models.Url, error)
	GetByUrl(ctx context.Context, url string) (*models.Url, error)
}
