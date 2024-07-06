package main

import (
    "databale/sql"
    "log"
    "os"
    _ "github.com/lib/pq"
)

func main() {
    pgURL := os.Getenv("PGURL")
    if pgURL == "" {
        log.Fatal("PGURL empty")
    }

    db, err := sql.Open("postgres", pgURL)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }
}
