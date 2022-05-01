package services

import (
	"github.com/ZhuoYIZIA/url-shortener/constants"
	"github.com/ZhuoYIZIA/url-shortener/models"
	"github.com/ZhuoYIZIA/url-shortener/repositories"
	"github.com/teris-io/shortid"
	"time"
)

func ConvertToShortUrl(originalUrl string) (*string, error) {
	currentTime := time.Now().In(constants.TimeZone)
	shortId, err := shortIdGenerator()
	if err != nil {
		return nil, err
	}

	// store
	err = repositories.StoreUrl(models.Url{
		Id:          *shortId,
		OriginalUrl: originalUrl,
		CreateAt:    currentTime,
	})
	if err != nil {
		return nil, err
	}

	return shortId, nil
}

func shortIdGenerator() (*string, error) {
	// make a shortid generate
	shortIdGenerate, err := shortid.New(1, shortid.DefaultABC, 2342)
	if err != nil {
		// TODO: error handler
		return nil, err
	}

	// generate id
	shortId, err := shortIdGenerate.Generate()
	if err != nil {
		// TODO: error handler
		return nil, err
	}
	return &shortId, nil
}
