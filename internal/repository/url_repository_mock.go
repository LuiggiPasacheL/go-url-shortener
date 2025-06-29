package repository

import (
	"context"
	"fmt"
	"slices"

	"github.com/LuiggiPasacheL/go-url-shortener/internal/models"
)

type UrlRepositoryMock struct {
	memory []*models.Url
}

func NewUrlRepositoryMock() *UrlRepositoryMock {
	return &UrlRepositoryMock{
		memory: make([]*models.Url, 0),
	}
}

func (r UrlRepositoryMock) Create(ctx context.Context, url *models.Url) (*models.Url, error) {

	l := len(r.memory)
	index := l + 1
	
	newUrl := &models.Url{
		Id: &index,
		LongUrl: url.LongUrl,
		ShortUrl: url.ShortUrl,
	}

	r.memory = append(r.memory, newUrl)

	return newUrl, nil
}

func (r UrlRepositoryMock) GetById(ctx context.Context, id int) (*models.Url, error) {
	
	index := slices.IndexFunc(r.memory, func (u *models.Url) bool {
		return u.Id	!= nil && *u.Id == id
	})

	if index == -1 {
		return nil, fmt.Errorf("Url doesn't exists")
	}

	return r.memory[index], nil
}

func (r UrlRepositoryMock) GetByShortUrl(ctx context.Context, shortUrl string) (*models.Url, error) {

	index := slices.IndexFunc(r.memory, func (u *models.Url) bool {
		return u.ShortUrl == shortUrl
	})

	if index == -1 {
		return nil, fmt.Errorf("Url doesn't exists")
	}

	return r.memory[index], nil
}

func (r UrlRepositoryMock) GetAll(ctx context.Context) ([]*models.Url, error) {
	return r.memory, nil
}

func (r UrlRepositoryMock) GetByUrl(ctx context.Context, url string) (*models.Url, error) {

	index := slices.IndexFunc(r.memory, func (u *models.Url) bool {
		return u.LongUrl == url
	})

	if index == -1 {
		return nil, fmt.Errorf("Url doesn't exists")
	}

	return r.memory[index], nil
}
