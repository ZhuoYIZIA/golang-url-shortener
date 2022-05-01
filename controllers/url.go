package controllers

import (
	"github.com/ZhuoYIZIA/url-shortener/request"
	"github.com/ZhuoYIZIA/url-shortener/services"
	"github.com/gin-gonic/gin"
)

// func GetUrl(c *gin.Context) {
// 	shortId := c.Param("id")

// }

func CreateUrl(c *gin.Context) {
	// get request data
	requestData := request.UrlRequest{}
	err := c.Bind(&requestData)
	if err != nil {
		// TODO: error handler
		c.JSON(400, err.Error())
		return
	}

	// validator
	err = requestData.Validator()
	if err != nil {
		// TODO: error handler
		c.JSON(400, err.Error())
		return
	}

	// convert to short url
	shortUrl, err := services.ConvertToShortUrl(requestData.Url)
	if err != nil {
		// TODO: error handler
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, shortUrl)
	return
}
