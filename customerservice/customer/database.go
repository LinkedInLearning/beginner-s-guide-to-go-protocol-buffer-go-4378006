package customer

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func InitializeDB() *sql.DB {
	os.Remove("./sqlcustomer.db")

	log.Println("Creating sqlcustomer.db...")
	file, err := os.Create("./sqlcustomer.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("sqlcustomer.db created")

	sqlDatabase, _ := sql.Open("sqlite3", "./sqlcustomer.db")
	defer sqlDatabase.Close()
	createTable(sqlDatabase)

	return sqlDatabase
}

func createTable(db *sql.DB) {
	log.Println("Createing table...")

	_, err := db.Exec("CREATE TABLE customers (id integer NOT NULL PRIMARY KEY AUTOINCREMENT,	username TEXT NOT NULL, password TEXT NOT NULL, email TEXT NOT NULL	)")

	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Table created")
}