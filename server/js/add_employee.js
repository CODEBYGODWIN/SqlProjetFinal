const postesParDepartement = {
    1: [
        { value: '1', text: 'PDG' },
        { value: '2', text: 'DGA' },
        { value: '3', text: 'Responsable Stratégie' }
    ],
    2: [
        { value: '4', text: 'Chef Comptable' },
        { value: '5', text: 'Comptable Fournisseurs/Clients' },
        { value: '6', text: 'Assistant Comptable' },
        { value: '7', text: 'Auditeur Interne' }
    ],
    3: [
        { value: '8', text: 'Directeur de la Recherche et Développement' },
        { value: '9', text: 'Parfumeur (Nez)' },
        { value: '10', text: 'Chimiste Formulateur' },
        { value: '11', text: 'Technicien de Laboratoire' }
    ],
    4: [
        { value: '12', text: 'Directeur Marketing' },
        { value: '13', text: 'Community Manager' }
    ],
    5: [
        { value: '14', text: 'Directeur Commercial' },
        { value: '15', text: 'Responsable des Ventes' }
    ],
    6: [
        { value: '16', text: 'Directeur de Production' },
        { value: '17', text: 'Responsable de la Chaîne de Production' }
    ],
    7: [
        { value: '18', text: 'Responsable Qualité' },
        { value: '19', text: 'Responsable des Conformités et Réglementations' }
    ],
    8: [
        { value: '20', text: 'Responsable Achats' },
        { value: '21', text: 'Gestionnaire des Stocks' }
    ],
    9: [
        { value: '22', text: 'Directeur Logistique' },
        { value: '23', text: 'Chef de Parc' },
        { value: '24', text: 'Transporteur / Livreur' },
        { value: '25', text: 'Planificateur Logistique' }
    ],
    10: [
        { value: '26', text: 'Directeur des Ressources Humaines' },
        { value: '27', text: 'Chargé de Recrutement' },
        { value: '28', text: 'Assistant RH' },
        { value: '29', text: 'Gestionnaire de Paie' }
    ],
    11: [
        { value: '30', text: 'Responsable du Service Client' },
        { value: '31', text: 'Chargé de Relations Clients' },
        { value: '32', text: 'Agent de Service Après-Vente (SAV)' },
        { value: '33', text: 'Gestionnaire de Réclamations' }
    ]
};

document.getElementById('departement').addEventListener('change', updatePosts);

function updatePosts() {
    const departementSelect = document.getElementById('departement');
    const postSelect = document.getElementById('post');
    const selectedDepartement = departementSelect.value;

    postSelect.innerHTML = '<option value="">Sélectionnez un poste</option>';

    if (postesParDepartement[selectedDepartement]) {
        postesParDepartement[selectedDepartement].forEach(function(post) {
            const option = document.createElement('option');
            option.value = post.value;
            option.text = post.text;
            postSelect.appendChild(option);
        });
    }
}

document.querySelector('form').addEventListener('submit', function(event) {
    event.preventDefault();
    
    const formData = new FormData(event.target);
    const data = {};

    formData.forEach((value, key) => {
        data[key] = value;
    });
    fetch('/employees/create', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
    })
    .then(response => response.text())
    .then(data => {
        alert('Employee created successfully!');
        console.log('Success:', data);
    })
    .catch((error) => {
        alert('Error creating employee.');
        console.error('Error:', error);
    });
});