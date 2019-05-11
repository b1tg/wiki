package server

import (
	"fmt"
	"net/http"

	"github.com/peterhellberg/wiki/db"
)

func (s *Server) show(w http.ResponseWriter, r *http.Request) {
	s.db.View(func(tx *db.Tx) error {
		p, err := tx.Page(s.getPageName(r))

		if err != nil || len(p.Text) == 0 {
			p.Text = emptyPageText
		}

		return show.Execute(w, data{
			"Title": string(p.Name),
			"Path":  "/" + string(p.Name) + "/edit",
			"Text":  bytesAsHTML(parsedMarkdown(p.Text)),
		})
	})
}

func (s *Server) list(w http.ResponseWriter, r *http.Request) {
	s.db.View(func(tx *db.Tx) error {
		names, _ := tx.Pages()
		fmt.Println("====got list: ", names)

		return list.Execute(w, names)
		// p, err := tx.Page(s.getPageName(r))

		// if err != nil || len(p.Text) == 0 {
		// 	p.Text = emptyPageText
		// }

		// return show.Execute(w, data{
		// 	"Title": string(p.Name),
		// 	"Path":  "/" + string(p.Name) + "/edit",
		// 	"Text":  bytesAsHTML(parsedMarkdown(p.Text)),
		// })
	})
}
