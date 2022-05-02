package repositories

import (
	"database/sql"
	"fmt"
	"github.com/ZhuoYIZIA/url-shortener/database"
	"log"
	"time"
)

type Url struct {
	Id          string    `json:"id"`
	OriginalUrl string    `json:"original_url"`
	CreateAt    time.Time `json:"create_at"`
}

func (url *Url) GetOriginalUrlInDB() string {
	db, err := database.PostgreSQLConnect()
	if err != nil {
		// TODO: error handler
		log.Fatal(err)
	}

	row := db.QueryRow("SELECT original_url FROM urls WHERE id = $1 LIMIT 1", url.Id)

	err = row.Scan(&url.OriginalUrl)
	if err != nil {
		// TODO: error handler
		log.Fatal(err)
	}

	return url.OriginalUrl
}

func (url *Url) StoreUrlInDB() error {
	db, err := database.PostgreSQLConnect()
	if err != nil {
		// TODO: error handler
		log.Fatal(err)
		//return err
	}

	stmt, err := db.Prepare("INSERT INTO urls(id, original_url, create_at) values ($1, $2, $3)")
	if err != nil {
		// TODO: error handler
		log.Fatal(err)
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			// TODO: error handler
			log.Fatal(err)
		}
	}(stmt)

	fmt.Println(url.Id)
	fmt.Println(url.OriginalUrl)

	_, err = stmt.Exec(url.Id, url.OriginalUrl, url.CreateAt)
	if err != nil {
		// TODO: error handler
		log.Fatal(err)
	}
	return nil
}
