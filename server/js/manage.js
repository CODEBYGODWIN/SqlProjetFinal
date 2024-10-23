document.getElementById('searchBtn').addEventListener('click', function () {
    const searchValue = document.getElementById('searchInput').value;
    const searchType = document.getElementById('searchType').value;

    fetch(`/employees/search?query=${searchValue}&type=${searchType}`)
        .then(response => response.json())
        .then(data => {
            const tbody = document.querySelector('#employeesTable tbody');
            tbody.innerHTML = '';
            data.forEach(employee => {
                const row = document.createElement('tr');
                row.innerHTML = `
                    <td>${employee.first_name}</td>
                    <td>${employee.last_name}</td>
                    <td>${employee.mail}</td>
                    <td>${employee.department}</td>
                    <td>${employee.post}</td>
                    <td>
                        <button onclick="editEmployee(${employee.id})">modifier</button>
                        <button onclick="deleteEmployee(${employee.id})">supprimer</button>
                    </td>
                `;
                tbody.appendChild(row);
            });
        });
});

function editEmployee(employeeId) {
    fetch(`/employeeid/${employeeId}`)
        .then(response => response.json())
        .then(employee => {
            document.getElementById('employeeId').value = employee.id;
            document.getElementById('firstName').value = employee.first_name;
            document.getElementById('lastName').value = employee.last_name;
            document.getElementById('mail').value = employee.mail;
            document.getElementById('address').value = employee.address;
            document.getElementById('phone').value = employee.phone;
            document.getElementById('salary').value = employee.salary;
            
            document.getElementById('employeeModal').style.display = 'block';
        });
}

document.querySelector('.close').onclick = function () {
    document.getElementById('employeeModal').style.display = 'none';
};

document.getElementById('employeeForm').addEventListener('submit', function (event) {
    event.preventDefault();
    
    const employeeId = document.getElementById('employeeId').value;
    const updatedEmployee = {
        first_name: document.getElementById('firstName').value,
        last_name: document.getElementById('lastName').value,
        mail: document.getElementById('mail').value,
        Address: document.getElementById('address').value,
        Phone: document.getElementById('phone').value,
        salary: document.getElementById('salary').value,
    };
    
    fetch(`/employees/${employeeId}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(updatedEmployee),
    })
    .then(response => response.json())
    .then(data => {
        if (data.message) {
            alert(data.message);
            document.getElementById('employeeModal').style.display = 'none';
            location.reload();
        }
    })
    .catch(error => {
        console.error('Error:', error);
        alert('Failed to update employee: ' + error.message);
    });
});

function deleteEmployee(employeeId) {
    if (confirm("Are you sure you want to delete this employee?")) {
        fetch(`/employee/${employeeId}`, {
            method: 'DELETE'
        }).then(response => {
            if (!response.ok) {
                return response.text().then(text => { throw new Error(text); });
            }
            location.reload();
        }).catch(error => {
            console.error('Error:', error);
            alert('Failed to delete employee: ' + error.message);
        });
    }
}
