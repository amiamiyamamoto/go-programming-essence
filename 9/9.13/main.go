package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// fmt.Println("start")
	db, err := sql.Open("sqlite3", "./database.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	queries := []string{
		`CREATE TABLE IF NOT EXISTS authors(author_id TEXT, author TEXT, PRIMARY KEY(author_id))`,
		`CREATE TABLE IF NOT EXISTS contents(author_id TEXT, title_id TEXT, title TEXT,content TEXT, PRIMARY KEY(author_id, title_id))`,
		`CREATE VIRTUAL TABLE IF NOT EXISTS contents_fts USING fts4(words)`,
	}
	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal(err)
		}
	}
	//処理
	b, err := os.ReadFile("./ababababa.txt")
	if err != nil {
		log.Fatal(err)
	}
	content := string(b)

	res, err := db.Exec(`INSERT INTO contents(author_id, title_id, title, content) VALUES(?,?,?,?)`,
		"0000879",
		"14",
		"あばばばば",
		content,
	)
	if err != nil {
		log.Fatal(err)
	}
	docID, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`INSERT INTO authors(author_id, author) VALUES(?, ?)`, "0000879", "作者名")
	if err != nil {
		log.Fatal(err)
	}

	t, err := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
	if err != nil {
		log.Fatal(err)
	}

	seg := t.Wakati(content)
	_, err = db.Exec(`
	INSERT INTO contents_fts(docid, words) VALUES(?,?)
	`,
		docID,
		strings.Join(seg, " "),
	)
	if err != nil {
		log.Fatal(err)
	}

	//contents_ftsテーブルに何件のデータが入っているか確認
	// ans, err := db.Query(`SELECT COUNT(*) FROM contents_fts`)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer ans.Close()
	// for ans.Next() {
	// 	var count int
	// 	err = ans.Scan(&count)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(count)
	// }

	//クエリを実行する
	query := "虫 AND ココア"
	//ここで発行されるクエリはこのようになっている
	// SELECT	a.author, c.title FROM contents c INNER JOIN authors a ON a.author_id = c.author_id INNER JOIN contents_fts f ON c.rowid = f.docid AND words MATCH '虫 AND ココア'
	//これと同じ条件で、取得されたデータの件数を確認する
	// SELECT COUNT(*) FROM contents c INNER JOIN authors a ON a.author_id = c.author_id INNER JOIN contents_fts f ON c.rowid = f.docid AND words MATCH '虫 AND ココア'
	// rows, err := db.Query(`
	// SELECT COUNT(*)
	// FROM
	// 	contents c
	// INNER JOIN authors a
	// 	ON a.author_id = c.author_id
	// INNER JOIN contents_fts f
	// 	ON c.rowid = f.docid
	// 	AND words MATCH ?
	// `, query)
	rows, err := db.Query(`
	SELECT
		a.author,
		c.title
	FROM
		contents c
	INNER JOIN authors a
		ON a.author_id = c.author_id
	INNER JOIN contents_fts f
		ON c.rowid = f.docid
		AND words MATCH ?
	`, query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		fmt.Println("rows")
		var author, title string
		// var count int
		err = rows.Scan(&author, &title)
		// err = rows.Scan(&count)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(author, title)
		// fmt.Println(count, "count")
	}
}
