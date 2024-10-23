package essencia

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Employee struct {
	ID         int    `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Mail       string `json:"mail"`
	Address    string `json:"address"`
	Phone      string `json:"phone"`
	HireDate   string `json:"hire_date"`
	Salary     string `json:"salary"`
	Department string `json:"department"`
	Post       string `json:"post"`
}

func DbConn() (db *sql.DB) {
	db, err := sql.Open("sqlite3", "./db/entreprise.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	tmplPath := fmt.Sprintf("templates/%s", tmpl)
	tmplParsed, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tmplParsed.Execute(w, nil)
}
