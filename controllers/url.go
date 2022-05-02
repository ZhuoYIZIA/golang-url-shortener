package controllers

import (
	"github.com/ZhuoYIZIA/url-shortener/request"
	"github.com/ZhuoYIZIA/url-shortener/services"
	"github.com/gin-gonic/gin"
)

func GetUrl(c *gin.Context) {
	shortId := c.Param("id")
	requestData := request.GetUrlRequest{Id: shortId}
	err := requestData.Validator()
	if err != nil {
		// TODO: error handler
		c.JSON(400, err.Error())
		return
	}

	originalUrl, err := services.GetOriginalUrlFromShortId(requestData.Id)
	if err != nil {
		// TODO: error handler
		c.JSON(400, err.Error())
		return
	}

	c.JSON(200, originalUrl)
	return
}

func CreateUrl(c *gin.Context) {
	// get request data
	requestData := request.CreateUrlRequest{}
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
