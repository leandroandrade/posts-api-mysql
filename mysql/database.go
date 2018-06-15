package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func init() {
	var err error

	DB, err = sql.Open("mysql", "root:root@tcp(container_mysql:3306)/golang_posts")
	if err != nil {
		log.Fatal(err.Error())
	} else {
		log.Print("Connection with database successful")
	}

	//err = DB.Ping()
	//if err != nil {
	//	log.Fatal("Cannot ping the database")
	//} else {
	//	log.Print("Database ping successful")
	//}

}
