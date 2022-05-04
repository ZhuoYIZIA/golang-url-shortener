package repositories

import (
	"context"
	"github.com/ZhuoYIZIA/url-shortener/constants"
	"github.com/ZhuoYIZIA/url-shortener/database"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	ctx                              = context.Background()
	rdb                              = database.RedisConnect()
	getUrlAndRefreshExpireTimeScript = redis.NewScript(`
		local key, ExpireTime = KEYS[1], ARGV[1]
		redis.call('EXPIRE', key, ExpireTime)
		local originalUrl = redis.call('GET', key)
		return originalUrl
	`)
)

type RedisUrl struct {
	Id          string    `json:"id"`
	OriginalUrl string    `json:"original_url"`
	CreateAt    time.Time `json:"create_at"`
}

func (url *RedisUrl) GetUrl() *string {
	keys := []string{url.Id}
	result, err := getUrlAndRefreshExpireTimeScript.Run(ctx, rdb, keys, constants.RedisExpireTime).Result()
	if err != nil {
		// TODO: error handle
	}
	if result == nil {
		return nil
	} else {
		originalUrl := result.(string)
		return &originalUrl
	}
}

func (url *RedisUrl) StoreUrl() error {
	err := rdb.Set(ctx, url.Id, url.OriginalUrl, constants.RedisExpireTime).Err()
	if err != nil {
		return err
	}
	return nil
}
