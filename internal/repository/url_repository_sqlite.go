package repository

import (
	"context"
	"database/sql"

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

	_, err = db.Exec("IF NOT EXISTS CREATE TABLE urls(id INTEGER PRIMARY KEY AUTOINCREMENT, url VARCHAR(250), shortUrl VARCHAR(250))")
	if err != nil {
		return nil, err
	}

	return repository, nil
}

func (r *UrlRepositorySqlite) Create(ctx context.Context, url *models.Url) (*models.Url, error) {
	stmt, err := r.db.PrepareContext(ctx, "INSERT INTO urls(url, shortUrl) VALUES (?, ?) ")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(url.Url, url.ShortUrl)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	id64 := int(id)

	newUrl := models.Url{
		Id: &id64,
		Url: url.Url,
		ShortUrl: url.ShortUrl,
	}

	return &newUrl, nil
}

func (r *UrlRepositorySqlite) GetById(ctx context.Context, id int) (*models.Url, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, url, shortUrl FROM urls WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	var sId int
	var sUrl, sShortUrl string
	for rows.Next() {
		err = rows.Scan(&sId, &sUrl, &sShortUrl)

		url := models.Url{
			Id: &sId,
			Url: sUrl,
			ShortUrl: sShortUrl,
		}

		return &url, nil
	}

	return nil, sql.ErrNoRows
}

func (r *UrlRepositorySqlite) GetByShortUrl(ctx context.Context, shortUrl string) (*models.Url, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, url, shortUrl FROM urls WHERE shortUrl = ?", shortUrl)
	if err != nil {
		return nil, err
	}

	var sId int
	var sUrl, sShortUrl string
	for rows.Next() {
		err = rows.Scan(&sId, &sUrl, &sShortUrl)

		url := models.Url{
			Id: &sId,
			Url: sUrl,
			ShortUrl: sShortUrl,
		}

		return &url, nil
	}

	return nil, sql.ErrNoRows
}

func (r *UrlRepositorySqlite) GetAll(ctx context.Context) ([]*models.Url, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, url, shortUrl FROM urls")
	if err != nil {
		return nil, err
	}

	var urls []*models.Url
	for rows.Next() {

		var sId int
		var sUrl, sShortUrl string
		err = rows.Scan(&sId, &sUrl, &sShortUrl)
		if err != nil {
			return nil, err
		}

		url := models.Url{
			Id: &sId,
			Url: sUrl,
			ShortUrl: sShortUrl,
		}

		urls = append(urls, &url)
	}
	if rows.Err() != nil {
		return nil, err
	}

	return urls, err
}


func (r *UrlRepositorySqlite) GetByUrl(ctx context.Context, url string) (*models.Url, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, url, shortUrl FROM urls WHERE url = ?", url)
	if err != nil {
		return nil, err
	}

	var sId int
	var sUrl, sShortUrl string
	for rows.Next() {
		err = rows.Scan(&sId, &sUrl, &sShortUrl)

		url := models.Url{
			Id: &sId,
			Url: sUrl,
			ShortUrl: sShortUrl,
		}

		return &url, nil
	}

	return nil, sql.ErrNoRows
}

func (r *UrlRepositorySqlite) Close() error {
	if r.db != nil {
		return r.db.Close()
	}
	return nil
}
