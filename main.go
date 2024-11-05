package main

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/devxbr/go-to-fulltack/microfrontends/card"
	"github.com/devxbr/go-to-fulltack/microfrontends/footer"
	"github.com/devxbr/go-to-fulltack/microfrontends/header"
	todo "github.com/devxbr/go-to-fulltack/microfrontends/to-do"
)

var startTime time.Time

func mainRender(w http.ResponseWriter) {
	tmpl, err := template.ParseFiles("templates/base.html")
	if err != nil {
		http.Error(w, "Erro ao carregar template base", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Header":   template.HTML(header.Render()),
		"TodoList": template.HTML(todo.Render()),
		"Card":     template.HTML(card.Render(startTime)),
		"Footer":   template.HTML(footer.Render()),
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Erro ao executar template base", http.StatusInternalServerError)
		return
	}
}
func main() {
	startTime = time.Now()
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mainRender(w)
	})

	http.HandleFunc("/add-to-do", func(w http.ResponseWriter, r *http.Request) {

		// Configura os cabeçalhos CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")              // Permite todas as origens
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS") // Permite métodos POST e OPTIONS
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")  // Permite cabeçalhos específicos

		// Verifica se a requisição é do tipo OPTIONS (necessário para CORS preflight)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Processa o POST normalmente
		if r.Method == http.MethodPost {
			var t todo.Todo
			f, _ := io.ReadAll(r.Body)
			err := json.Unmarshal(f, &t)
			if err != nil {
				http.Error(w, "Erro ao ler corpo da requisição", http.StatusBadRequest)
				return
			}
			todo.Add(t)
			http.Redirect(w, r, "/", http.StatusOK)
			return
		} else {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
