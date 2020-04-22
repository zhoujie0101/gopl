package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/sakila")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query(db)

	prepare(db)
}

func query(db *sql.DB) {
	rows, err := db.Query("select film_id, title from film where film_id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var filmID, title string
	for rows.Next() {
		err = rows.Scan(&filmID, &title)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(filmID, title)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func prepare(db *sql.DB) {
	stmt, err := db.Prepare("select film_id, title from film where film_id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(1)
	defer rows.Close()
	var filmID, title string
	for rows.Next() {
		err = rows.Scan(&filmID, &title)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(filmID, title)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
