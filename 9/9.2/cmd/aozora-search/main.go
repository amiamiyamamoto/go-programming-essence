package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var dsn string
	flag.StringVar(&dsn, "d", "database.sqlite", "database")
	usage := `
	Usage:
	  -d string       specify the database file (default "database.sqlite")
	  
	Commands:
	  authors         show all authors
	  titles          show titles for a specific author
	  content         show content for a specific author and title
	  query           perform a query on the content
	
	Flags:
	  -h              show help message
	`
	flag.Usage = func() {
		fmt.Print(usage)
	}
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(2)
	}

	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	switch flag.Arg(0) {
	case "authors":
		err = showAuthors(db)
	case "titles":
		if flag.NArg() != 2 {
			flag.Usage()
			os.Exit(2)
		}
		err = showTitles(db, flag.Arg(1))
	case "content":
		if flag.NArg() != 3 {
			flag.Usage()
			os.Exit(2)
		}
		err = showContent(db, flag.Arg(1), flag.Arg(2))
	case "query":
		if flag.NArg() != 2 {
			flag.Usage()
			os.Exit(2)
		}
		err = queryContent(db, flag.Arg(1))
	}
	if err != nil {
		log.Fatal(err)
	}
}
func showAuthors(db *sql.DB) error {
	return nil
}
func showTitles(db *sql.DB, authorID string) error {
	return nil
}
func showContent(db *sql.DB, authorID string, titleID string) error {
	return nil
}
func queryContent(db *sql.DB, query string) error {
	return nil
}
