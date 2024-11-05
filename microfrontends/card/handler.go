package card

import (
	"fmt"
	"time"
)

type Card struct {
	ImageURL string
	Title    string
	Desc     string
}

func Render(startTime time.Time) string {
	uptime := time.Since(startTime)
	card := Card{
		ImageURL: "/static/assets/gopher_nerd.png",
		Title:    "Informação do servidor",
		Desc:     fmt.Sprintf("Server uptime: %s", uptime),
	}
	return fmt.Sprintf(`
            <div class="card">
                <div class="card-image">
                    <img src="%s" alt="%s">
                    <span class="card-title">%s</span>
                </div>
                <div class="card-content">
                    <p>%s</p>
                </div>
            </div>
    `, card.ImageURL, card.Title, card.Title, card.Desc)
}
