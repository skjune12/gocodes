package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var dbfile = "./test.db"
	os.Remove(dbfile)
	db, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE "world" (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"country" VARCHAR(255),
		"capital" VARCHAR(255)
	)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`INSERT INTO "world" ("country", "capital") VALUES(?, ?)`, "Japan", "Tokyo")

	if err != nil {
		log.Fatal(err)
	}

	// prepare() で宣言して、Close() で閉じる
	stmt, err := db.Prepare(`INSERT INTO "world" ("country", "capital") VALUES (?, ?)`)
	if err != nil {
		log.Fatal(err)
	}

	if _, err = stmt.Exec("America", "Washington, D.C."); err != nil {
		log.Fatal(err)
	}

	if _, err = stmt.Exec("Russia", "Moscow"); err != nil {
		log.Fatal(err)
	}

	if _, err = stmt.Exec("British", "London"); err != nil {
		log.Fatal(err)
	}

	if _, err = stmt.Exec("Australia", "Sydney"); err != nil {
		log.Fatal(err)
	}

	stmt.Close()
	db.Close()
}
