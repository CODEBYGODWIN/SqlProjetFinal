package essencia

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	db := DbConn()
	defer db.Close()

	rows, err := db.Query(`SELECT Employees.employeeId, firstName, lastName, mail, address, phone, hireDate, salary, Departements.name, Posts.postName 
		FROM Employees 
		JOIN Departements ON Employees.departementId = Departements.departementId
		JOIN Posts ON Employees.postId = Posts.postId`)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var employees []Employee
	for rows.Next() {
		var emp Employee
		err := rows.Scan(&emp.ID, &emp.FirstName, &emp.LastName, &emp.Mail, &emp.Address, &emp.Phone, &emp.HireDate, &emp.Salary, &emp.Department, &emp.Post)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		employees = append(employees, emp)
	}

	json.NewEncoder(w).Encode(employees)
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	db := DbConn()
	defer db.Close()

	var emp Employee
	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	add, err := db.Prepare("INSERT INTO Employees (firstName, lastName, mail, address, phone, hireDate, salary, departementId, postId) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	_, err = add.Exec(emp.FirstName, emp.LastName, emp.Mail, emp.Address, emp.Phone, emp.HireDate, emp.Salary, emp.Department, emp.Post)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, "Employee created successfully")
}
