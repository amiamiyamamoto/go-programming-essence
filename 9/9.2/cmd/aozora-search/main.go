package main

import (
	"archive/zip"
	"bytes"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/text/encoding/japanese"
)

type Entry struct {
	AuthorID string
	Author   string
	TitleID  string
	Title    string
	SiteURL  string
	ZipURL   string
}

// findEntries は指定されたURLから作品一覧を取得する
func findEntries(siteURL string) ([]Entry, error) {
	res, err := http.Get(siteURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("Failed to fetch page: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}
	pat := regexp.MustCompile(`.*/cards/([0-9]+)/card([0-9]+).html$`)
	entiries := []Entry{}
	doc.Find("ol li a").Each(func(n int, elem *goquery.Selection) {
		token := pat.FindStringSubmatch(elem.AttrOr("href", ""))
		if len(token) != 3 {
			return
		}
		title := elem.Text()
		pageURL := fmt.Sprintf("https://www.aozora.gr.jp/cards/%s/card%s.html", token[1], token[2])
		author, zipURL := findAuthorAndZIP(pageURL) // 作者とZIPファイルのURLを得る
		if zipURL != "" {
			entiries = append(entiries, Entry{
				AuthorID: token[1],
				Author:   author,
				TitleID:  token[2],
				Title:    title,
				SiteURL:  siteURL,
				ZipURL:   zipURL,
			})
		}
		println(zipURL)
	})
	return entiries, nil
}

// 作者とZIPファイルのURLを得る
func findAuthorAndZIP(siteURL string) (string, string) {
	res, err := http.Get(siteURL)
	if err != nil {
		return "", ""
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("Failed to fetch page: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", ""
	}
	author := doc.Find("table[summary=作家データ] tr:nth-child(1) td:nth-child(2)").Text()

	zipURL := ""
	doc.Find("table.download a").Each(func(n int, elem *goquery.Selection) {
		href := elem.AttrOr("href", "")
		if strings.HasSuffix(href, ".zip") {
			zipURL = href
		}
	})
	u, err := url.Parse(siteURL)
	if err != nil {
		return author, ""
	}
	u.Path = path.Join(path.Dir(u.Path), zipURL)
	return author, u.String()
}

// ZIPファイルからテキストを抽出する
func extractText(zipURL string) (string, error) {
	resp, err := http.Get(zipURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	r, err := zip.NewReader(bytes.NewReader(b), int64(len(b)))
	if err != nil {
		return "", err
	}

	for _, file := range r.File {
		if path.Ext(file.Name) == ".txt" {
			f, err := file.Open()
			if err != nil {
				return "", err
			}
			b, err := io.ReadAll(f)
			f.Close()
			if err != nil {
				return "", err
			}
			b, err = japanese.ShiftJIS.NewDecoder().Bytes(b)
			if err != nil {
				return "", err
			}
			return string(b), nil
		}
	}
	return "", fmt.Errorf("text file not found in %s", zipURL)
}

// 作品リストをデータベースに登録する
func setupDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	queries := []string{
		`CREATE TABLE IF NOT EXISTS authors(author_id TEXT, author TEXT, PRIMARY KEY(author_id))`,
		`CREATE TABLE IF NOT EXISTS contents(author_id TEXT, title_id TEXT, title TEXT, cotent TEXT, PRIMARY KEY(author_id, title_id))`,
		`CREATE VIRTUAL TABLE IF NOT EXISTS contents_fts USING fts4(words)`,
	}
	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}

func addentry(db *sql.DB, entry *Entry, content string) error {
	_, err := db.Exec(`
		REPLACE INTO authors(author_id, author) VALUES(?, ?)
	`,
		entry.AuthorID,
		entry.Author,
	)
	if err != nil {
		return err
	}

	res, err := db.Exec(`
		REPLACE INTO contents(author_id, title_id, title, content) VALUES(?,?,?,?)
		`,
		entry.AuthorID,
		entry.TitleID,
		entry.Title,
		content,
	)
	if err != nil {
		return err
	}
	docID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	t, err := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
	if err != nil {
		return err
	}

	seg := t.Wakati(content)
	_, err = db.Exec(`
	RELPACE INTO contents_fts(docid, words) VALUES(?,?)
	`,
		docID,
		strings.Join(seg, " "),
	)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	db, err := setupDB("database.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	listURL := "https://www.aozora.gr.jp/index_pages/person879.html"

	entries, err := findEntries(listURL)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("found %d entries", len(entries))
	for _, entry := range entries {
		content, err := extractText(entry.ZipURL)
		if err != nil {
			log.Println(err)
			continue
		}
		err = addentry(db, &entry, content)
		if err != nil {
			log.Println(err)
			continue
		}
		// fmt.Println(entry.Title, entry.ZipURL)
		// fmt.Println(content)
	}
}
