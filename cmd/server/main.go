package main

import (
	"crispy-website/internal/db"
	"crispy-website/internal/repo"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {
	repository, err := repo.NewRepository("walletdrain.db")
	if err != nil {
		panic("failed to open db")
	}

	fs := http.FileServer(http.Dir("static"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/index.html")
	})

	http.HandleFunc("/music", func(w http.ResponseWriter, r *http.Request) {
		results := repository.LoadMusic()
		spending := repository.GetSpending()
		tmpl := template.Must(template.ParseFiles("templates/music.html"))

		data := db.MusicData{Releases: results, Spending: spending}

		err := tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/.well-known/discord", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("dh=b81129b4e9cc388c5ab63919550316fc3ca5ebe4"))
	})

	fmt.Println("Running on :8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
