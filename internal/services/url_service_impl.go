package services

import (
	"context"
	"errors"

	"github.com/LuiggiPasacheL/go-url-shortener/internal/models"
	"github.com/LuiggiPasacheL/go-url-shortener/internal/repository"
)

type UrlServiceImpl struct {
	repository repository.UrlRepository
}

func NewUrlServiceImpl(repository repository.UrlRepository) *UrlServiceImpl {
	return &UrlServiceImpl{
		repository: repository,
	}
}

func (s *UrlServiceImpl) CreateUrl(ctx context.Context, url string) (*models.Url, error) {
	u, err := s.repository.GetByUrl(ctx, url)
	if err != nil {
		return nil, err
	}

	if u != nil {
		return nil, errors.New("Error url exists")
	}

	newUrl := models.Url{
		Url: url,
	}

	u, err = s.repository.Create(ctx, &newUrl)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *UrlServiceImpl) GetUrl(ctx context.Context, id int) (*models.Url, error) {
	u, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *UrlServiceImpl) GetAllUrls(ctx context.Context) ([]*models.Url, error) {
	urls, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return urls, nil
}

func (s *UrlServiceImpl) RedirectUrl(ctx context.Context, shortUrl string) (*models.Url, error) {
	u, err := s.repository.GetByShortUrl(ctx, shortUrl)
	if err != nil {
		return nil, err
	}

	return u, nil
}
