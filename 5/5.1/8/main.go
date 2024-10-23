package main

import (
	"net/http"
	"path"
)

func main() {
	mimemap := map[string]string{
		".xls":  "application/vnd.ms-excel",
		".xlsx": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		".ppt":  "application/vnd.ms-powerpoint",
		".pptx": "application/vnd.openxmlformats-officedocument.presentationml.presentation",
		".doc":  "application/msword",
		".docx": "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
	}
	fileserver := http.StripPrefix("/public/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// if path.Ext(r.URL.Path) == ".xls" {
		// 	w.Header().Set("Content-Type", "application/vnd.ms-excel")
		// }
		if typ, ok := mimemap[path.Ext(r.URL.Path)]; ok {
			w.Header().Set("Content-Type", typ)
		}
		fileserver.ServeHTTP(w, r)
	})

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./static"))))
	http.ListenAndServe(":8080", nil)
}
