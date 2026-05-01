package db

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "log"
)

func SetupDb() *sql.DB {
    db, err := sql.Open("sqlite3", "rssreader.db")
    if err != nil {
        log.Fatal(err.Error())
    }
    if err = db.Ping(); err != nil {
        log.Fatal(err.Error())
    }
    return db
}
