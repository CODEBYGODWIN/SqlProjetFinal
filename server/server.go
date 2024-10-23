package main

import (
	essencia "essencia/server/functions"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	essencia.DbConn()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		essencia.RenderTemplate(w, "home.html")
	})
	http.HandleFunc("/add_employee", func(w http.ResponseWriter, r *http.Request) {
		essencia.RenderTemplate(w, "add_employee.html")
	})
	http.HandleFunc("/manage_employee", func(w http.ResponseWriter, r *http.Request) {
		essencia.RenderTemplate(w, "manage.html")
	})

	http.HandleFunc("/employees", essencia.GetEmployees)
	http.HandleFunc("/employeeid/{id}", essencia.GetEmployeeByID)
	http.HandleFunc("/employees/create", essencia.CreateEmployee)
	http.HandleFunc("/employees/search", essencia.SearchEmployees)
	http.HandleFunc("/employees/{id}", essencia.UpdateEmployee)
	http.HandleFunc("/employee/{id}", essencia.DeleteEmployee)

	fmt.Println("Server listening on http://localhost:4040/")
	log.Fatal(http.ListenAndServe(":4040", nil))
}
