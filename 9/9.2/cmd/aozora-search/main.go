package mai

import (
	"fmt"
	"log"
)

type Entry struct {
	AuthorID string
	Auther   string
	TitleID  string
	Title    string
	InfoURL  string
	ZipURL   string
}

func findrntries(siteURL string) ([]Entiry, error) {
	//処理
}

func main() {
	listURL := "https//www.aozora.gr.jp/index_pages/person879.html"

	entries, err := findEntries(lisURL)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		fmt.Println(entry.Title, entry.ZipURL)
	}
}
