package essencia

import (
	"encoding/json"
	"net/http"
)

func SearchEmployees(w http.ResponseWriter, r *http.Request) {

	db := DbConn()
	defer db.Close()

	query := r.URL.Query().Get("query")
	searchType := r.URL.Query().Get("type")

	var sqlQuery string

	switch searchType {
	case "name":
		sqlQuery = `SELECT employeeId, firstName, lastName, mail, address, phone, hireDate, salary, Departements.name, Posts.postName 
					FROM Employees 
					JOIN Departements ON Employees.departementId = Departements.departementId
					JOIN Posts ON Employees.postId = Posts.postId
					WHERE firstName LIKE ? OR lastName LIKE ?`
	case "post":
		sqlQuery = `SELECT employeeId, firstName, lastName, mail, address, phone, hireDate, salary, Departements.name, Posts.postName 
					FROM Employees 
					JOIN Departements ON Employees.departementId = Departements.departementId
					JOIN Posts ON Employees.postId = Posts.postId
					WHERE Posts.postName LIKE ?`
	case "department":
		sqlQuery = `SELECT employeeId, firstName, lastName, mail, address, phone, hireDate, salary, Departements.name, Posts.postName 
					FROM Employees 
					JOIN Departements ON Employees.departementId = Departements.departementId
					JOIN Posts ON Employees.postId = Posts.postId
					WHERE Departements.name LIKE ?`
	}

	rows, err := db.Query(sqlQuery, "%"+query+"%", "%"+query+"%")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var employees []Employee
	for rows.Next() {
		var emp Employee
		if err := rows.Scan(&emp.ID, &emp.FirstName, &emp.LastName, &emp.Mail, &emp.Address, &emp.Phone, &emp.HireDate, &emp.Salary, &emp.Department, &emp.Post); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		employees = append(employees, emp)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}
