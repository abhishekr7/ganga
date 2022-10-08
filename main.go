package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/lib/pq"
)

var db *sql.DB

func main() {
    // Capture connection properties.
    user   := os.Getenv("user")
    host   := os.Getenv("host")
    dbname := os.Getenv("dbname")
    port   := os.Getenv("port")
    password := os.Getenv("password")
    // Get a database handle.
    var err error

    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")
}

