package app

import (
	"c/Users/winst/go/Go_Examples/imagesUrl_06/internal/parser"
	"c/Users/winst/go/Go_Examples/imagesUrl_06/internal/store"
	m "c/Users/winst/go/Go_Examples/imagesUrl_06/model"
	"c/Users/winst/go/Go_Examples/imagesUrl_06/repo"
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//URLColors ...
var URLColors []m.URLImage = parser.GetImagesLinks()

var db *sql.DB = store.CreateDB()

//RunBD ...
func RunBD() {

	store.CreateTable(db)

	repo.InsertURL(db, URLColors)

	//defer db.Close()
}

//RunAPI ...
func RunAPI() {

	router := mux.NewRouter()
	log.Print("Router starting...")

	router.HandleFunc("/", handleURLImages(db))

	err := http.ListenAndServe(":8181", router)
	if err != nil {
		log.Print(err)
	} else {
		log.Printf("starting api server ...")
	}
}
