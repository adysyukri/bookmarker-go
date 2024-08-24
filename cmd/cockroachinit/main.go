package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var dbUri = "postgresql://root@localhost:26257/bookmark_go?sslmode=disable"

func main() {
	conn, err := pgx.Connect(context.Background(), dbUri)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	// uncomment if error init table
	// _, err = conn.Exec(context.Background(), "DROP DATABASE bookmark_go;")
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "drop database failed: %v\n", err)
	// 	os.Exit(1)
	// }

	_, err = conn.Exec(context.Background(), "CREATE DATABASE bookmark_go;")
	if err != nil {
		fmt.Fprintf(os.Stderr, "create database failed: %v\n", err)
		os.Exit(1)
	}

	sqlStmt := `
	CREATE TABLE bookmarks (
		id varchar(40) PRIMARY KEY,
		title varchar,
		author varchar,
		total INT,
		read INT,
		created_at TIMESTAMP
	);`

	_, err = conn.Exec(context.Background(), sqlStmt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "create table failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("init table success!")
}
