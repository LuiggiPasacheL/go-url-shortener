package models

type Url struct {
	Id       *int   `json:"id,omitempty"`
	Url      string `json:"url,omitempty"`
	ShortUrl string `json:"short_url,omitempty"`
}
