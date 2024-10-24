package essencia

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
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

func GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/employeeid/"):]

	db := DbConn()
	defer db.Close()

	var emp Employee
	query := `SELECT Employees.employeeId, firstName, lastName, mail, address, phone, hireDate, salary, Departements.name, Posts.postName 
	          FROM Employees 
	          JOIN Departements ON Employees.departementId = Departements.departementId
	          JOIN Posts ON Employees.postId = Posts.postId
	          WHERE Employees.employeeId = ?`

	err := db.QueryRow(query, id).Scan(&emp.ID, &emp.FirstName, &emp.LastName, &emp.Mail, &emp.Address, &emp.Phone, &emp.HireDate, &emp.Salary, &emp.Department, &emp.Post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(emp)
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

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	db := DbConn()
	defer db.Close()
	id := r.URL.Path[len("/employees/"):]

	var emp Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `UPDATE Employees SET firstName=?, lastName=?, mail=?, address=?, phone=?, salary=? WHERE employeeId=?`
	_, err := db.Exec(query, emp.FirstName, emp.LastName, emp.Mail, emp.Address, emp.Phone, emp.Salary, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Employee updated successfully"})
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	db := DbConn()
	defer db.Close()

	idStr := strings.TrimPrefix(r.URL.Path, "/employee/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	query := `DELETE FROM Employees WHERE employeeId = ?`
	_, err = db.Exec(query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Employee deleted successfully")
}
