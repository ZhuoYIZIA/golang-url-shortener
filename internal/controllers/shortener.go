package controllers

import (
	"fmt"
	"os"
)

func Convert() {
	fmt.Println("test")
	db_host := os.Getenv("DB_HOST")
	fmt.Println(db_host)
}