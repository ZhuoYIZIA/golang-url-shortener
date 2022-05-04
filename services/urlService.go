package services

import (
	"github.com/ZhuoYIZIA/url-shortener/constants"
	"github.com/ZhuoYIZIA/url-shortener/repositories"
	"time"
)

func GetOriginalUrlFromShortId(shortId string) (*string, error) {
	redisUrl := getOriginalUrlFromRedis(shortId)

	if redisUrl != nil {
		return redisUrl, nil
	} else {
		originalUrl := getOriginalUrlFromPostgre(shortId)
		return originalUrl, nil
	}
}

func getOriginalUrlFromRedis(shortId string) *string {
	redisUrlRepo := repositories.RedisUrl{Id: shortId}
	url := redisUrlRepo.GetUrl()
	return url
}

func getOriginalUrlFromPostgre(shortId string) *string {
	dbUrlRepo := repositories.PgUrl{Id: shortId}
	url := dbUrlRepo.GetOriginalUrl()
	return url
}

func ConvertToShortUrl(originalUrl string) (*string, error) {
	currentTime := time.Now().In(constants.TimeZone)
	shortId, err := shortIdGenerator()
	if err != nil {
		// TODO: error handler
		return nil, err
	}

	CacheUrl := repositories.RedisUrl{
		Id:          *shortId,
		OriginalUrl: originalUrl,
		CreateAt:    currentTime,
	}
	err = CacheUrl.StoreUrl()
	if err != nil {
		// TODO: error handler
		return nil, err
	}

	DBUrl := repositories.PgUrl{
		Id:          *shortId,
		OriginalUrl: originalUrl,
		CreateAt:    currentTime,
	}
	err = DBUrl.StoreUrl()
	if err != nil {
		// TODO: error handler
		return nil, err
	}

	return shortId, nil
}
