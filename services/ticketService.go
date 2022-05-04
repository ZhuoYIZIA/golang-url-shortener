package services

import (
	"github.com/teris-io/shortid"
)

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
