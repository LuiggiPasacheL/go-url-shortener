package dto

type Url struct {
	Id       *int   `json:"id,omitempty"`
	LongUrl  string `json:"longUrl,omitempty" binding:"required"`
	ShortUrl string `json:"shortUrl,omitempty"`
}
