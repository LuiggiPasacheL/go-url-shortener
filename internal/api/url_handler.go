package api

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/LuiggiPasacheL/go-url-shortener/internal/api/dto"
	"github.com/LuiggiPasacheL/go-url-shortener/internal/services"
	"github.com/LuiggiPasacheL/go-url-shortener/pkg/log"
	"github.com/gin-gonic/gin"
)

type UrlHandler struct {
    urlService services.UrlService
}

func NewUrlHandler(logger *slog.Logger, urlService services.UrlService) UrlHandler {
    return UrlHandler{
        urlService: urlService,
    }
}

func (h UrlHandler) GetAllUrls(c *gin.Context) {

    logger := log.GetLogger(c)

    urls, err := h.urlService.GetAllUrls(c)
    if err != nil {
        logger.Error("Failed to fetch URLs", "error", err)
        c.JSON(http.StatusInternalServerError, dto.BaseResponse{
            Code:  http.StatusInternalServerError,
            Message: "Failed to fetch URLs",
        })
        return
    }

    if len(urls) == 0 {
        logger.Info("No URLs found")
        c.JSON(http.StatusNotFound, dto.BaseResponse{
            Code:    http.StatusNotFound,
            Message: "No URLs found",
        })
        return
    }

    urlDtos := make([]dto.Url, len(urls))
    for id := range urls {
        urlDtos[id] = dto.Url{
            Id:       urls[id].Id,
            ShortUrl: urls[id].ShortUrl,
        }
    }

    c.JSON(http.StatusOK, dto.BaseResponse{
        Code:    http.StatusOK,
        Message: "URLs fetched successfully",
        Data:    urlDtos,
    })
}

func (h UrlHandler) CreateUrl(c *gin.Context) {

    logger := log.GetLogger(c)

    var urlDto dto.Url
    if err := c.ShouldBindJSON(&urlDto); err != nil {
        logger.Error("Failed to bind JSON", "error", err)
        c.JSON(http.StatusBadRequest, dto.BaseResponse{
            Code:    http.StatusBadRequest,
            Message: "Invalid input",
        })
        return
    }

    url, err := h.urlService.CreateUrl(c, urlDto.Url)
    if err != nil {
        logger.Error("Failed to create URL", "error", err)
        c.JSON(http.StatusInternalServerError, dto.BaseResponse{
            Code:    http.StatusInternalServerError,
            Message: "Failed to create URL",
        })
        return
    }

    c.JSON(http.StatusCreated, dto.BaseResponse{
        Code:    http.StatusCreated,
        Message: "URL created successfully",
        Data: dto.Url{
            Id:       url.Id,
            ShortUrl: url.ShortUrl,
        },
    })
}

func (h UrlHandler) RedirectUrl(c *gin.Context) {

    logger := log.GetLogger(c)

    shortUrl := c.Param("s")
    if shortUrl == "" {
        logger.Error("Short URL is empty")
        c.JSON(http.StatusBadRequest, dto.BaseResponse{
            Code:    http.StatusBadRequest,
            Message: "Short URL is required",
        })
        return
    }

    url, err := h.urlService.RedirectUrl(c, shortUrl)
    if err != nil {
        logger.Error("Failed to redirect URL", "error", err)
        c.JSON(http.StatusNotFound, dto.BaseResponse{
            Code: http.StatusNotFound,
            Message: "Failed to get url",
        })
        return
    }

    c.Redirect(http.StatusPermanentRedirect, url.ShortUrl)
    
}

func (h UrlHandler) GetUrl(c *gin.Context) {

    logger := log.GetLogger(c)

    idStr := c.Param("urlId")

    if idStr == "" {
        logger.Error("id is empty")
        c.JSON(http.StatusBadRequest, dto.BaseResponse{
            Code: http.StatusBadRequest,
            Message: "Id is required",
        })
        return
    }

    id, err := strconv.Atoi(idStr)
    if err != nil {
        logger.Error("id cannot be converted into int", "error", err)
        c.JSON(http.StatusBadRequest, dto.BaseResponse{
            Code: http.StatusBadRequest,
            Message: "Id must be an integer " + idStr,
        })
    }

    url, err := h.urlService.GetUrl(c, id)
    if err != nil {
        logger.Error("Failed to get Url", "error", err)
        c.JSON(http.StatusInternalServerError, dto.BaseResponse{
            Code: http.StatusInternalServerError,
            Message: "Failed to get url",
        })
    }

    urlDto := dto.Url{
        Id: url.Id,
        Url: url.Url,
        ShortUrl: url.ShortUrl,
    }

    c.JSON(http.StatusOK, dto.BaseResponse{
        Code: http.StatusInternalServerError,
        Message: "Ok",
        Data: urlDto,
    })
}
