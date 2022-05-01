package main

import (
	"github.com/ZhuoYIZIA/url-shortener/routers"
)

func main() {
	engine := routers.InitRoutes()
	panic(engine.Run())
}
