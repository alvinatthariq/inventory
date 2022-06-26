package main

import (
	"database/sql"
	"log"

	restserver "inventory/rest"
	"inventory/stock"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

var (
	db  *sql.DB
	err error
)

func main() {
	// init router
	router := httprouter.New()

	// init sql connection
	db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1)/inventory_db")
	if err != nil {
		log.Fatalln(err)
	}

	log.Print("ping mysql...")

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	log.Print("mysql ok")

	stockItf := stock.InitStock(db)

	restserver.Init(router, stockItf)

	defer func() {
		db.Close()
	}()
}
