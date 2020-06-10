package app

import (
	"c/Users/winst/go/Go_Examples/imagesUrl_06/model"
	"c/Users/winst/go/Go_Examples/imagesUrl_06/repo"
	"database/sql"
	"html/template"
	"log"
	"net/http"
)

func handleURLImages(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		d := repo.GetImg(db)

		data := model.Data{
			URLImgData: d,
		}

		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			log.Print(err)
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			log.Print(err)
		} else {
			log.Print("Ok...")
		}
	}
}
