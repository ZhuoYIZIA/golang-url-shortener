package repositories

import (
	"database/sql"
	"fmt"
	"github.com/ZhuoYIZIA/url-shortener/database"
	"github.com/ZhuoYIZIA/url-shortener/models"
	"log"
)

func StoreUrl(url models.Url) error {
	db, err := database.PostgreSQLConnect()
	if err != nil {
		// TODO: error handler
		log.Fatal(err)
		//return err
	}

	stmt, err := db.Prepare("INSERT INTO urls(id, original_url, create_at) values ($1, $2, $3)")
	if err != nil {
		log.Fatal(err)
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(stmt)

	fmt.Println(url.Id)
	fmt.Println(url.OriginalUrl)

	_, err = stmt.Exec(url.Id, url.OriginalUrl, url.CreateAt)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
