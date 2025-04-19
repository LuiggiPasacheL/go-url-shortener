package services

import (
	"context"

	"github.com/LuiggiPasacheL/go-url-shortener/internal/models"
)

type UrlService interface {
    CreateUrl(ctx context.Context, url string) (*models.Url, error)
    GetUrl(ctx context.Context, id int) (*models.Url, error)
    GetAllUrls(ctx context.Context) ([]*models.Url, error)
    RedirectUrl(ctx context.Context, shortUrl string) (*models.Url, error)
}


