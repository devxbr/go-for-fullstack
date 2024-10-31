package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/devxbr/go-for-fulltack/microfrontends/card"
	"github.com/devxbr/go-for-fulltack/microfrontends/footer"
	"github.com/devxbr/go-for-fulltack/microfrontends/header"
)

var startTime time.Time

func main() {
	startTime = time.Now()
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/base.html")
		if err != nil {
			http.Error(w, "Erro ao carregar template base", http.StatusInternalServerError)
			return
		}

		data := map[string]interface{}{
			"Header": template.HTML(header.Render()),
			"Card":   template.HTML(card.Render(startTime)),
			"Footer": template.HTML(footer.Render()),
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Erro ao executar template base", http.StatusInternalServerError)
			return
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
