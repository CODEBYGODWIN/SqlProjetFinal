let isEmployeeListVisible = false;

function fetchEmployees() {
    if (isEmployeeListVisible) {
        document.getElementById("employee-table").style.display = "none";
        isEmployeeListVisible = false;
        return;
    }
    fetch("/employees")
        .then(response => response.json())
        .then(data => {
            const employeeList = document.getElementById("employee-list");
            employeeList.innerHTML = '';
            data.forEach(employee => {
                const row = document.createElement("tr");
                row.innerHTML = `
                    <td>${employee.id}</td>
                    <td>${employee.first_name}</td>
                    <td>${employee.last_name}</td>
                    <td>${employee.mail}</td>
                    <td>${employee.address}</td>
                    <td>${employee.phone}</td>
                    <td>${employee.hire_date}</td>
                    <td>${employee.salary}</td>
                    <td>${employee.department}</td>
                    <td>${employee.post}</td>
                `;
                employeeList.appendChild(row);
            });
            document.getElementById("employee-table").style.display = "table";
            isEmployeeListVisible = true;
        })
        .catch(error => console.error('Erreur lors de la récupération des employés:', error));
}

document.getElementById("show-employees").addEventListener("click", fetchEmployees);
