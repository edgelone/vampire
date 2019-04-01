package data

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB

type DbWorker struct {
	DataResource string
}

func init()  {
	var err error
	Db,err =sql.Open("mysql","root:231@tcp(127.0.0.1:3306)/vampire")
	if err !=nil{
		log.Fatal(err)
	}
	return
}