CREATE TABLE Employees (

employeeId integer PRIMARY key not null,
firstName VARCHAR(50) not null,
lastName varchar(50) not NULL,
mail VARCHAR(50) not null,
address VARCHAR(100) not null,
phone VARCHAR(50) not null,
hireDate date not null,
salary INTEGER not null,
departementId INTEGER not null,
postId INTEGER not null,
FOREIGN KEY (departementId) REFERENCES Departements(departementId),
FOREIGN KEY (postId) REFERENCES Posts(postId)
);

CREATE TABLE Departements (
departementId integer PRIMARY key not null,
name VARCHAR(50) not null,
dPriority integer not null
);


CREATE TABLE Posts (

postId integer PRIMARY key not null,
postName VARCHAR(50) not null,
pPriority integer not null,
departementId INTEGER not null,
FOREIGN KEY (departementId) REFERENCES Departements(departementId)

)


INSERT INTO Departements (name, dPriority) VALUES 
('DIRECTION GENERALE', 1),
('COMPTABILITE', 2),
('MARKETING', 2),
('RECHERCHE ET DEVELOPPEMENT', 3),
('PRODUCTION/FABRICATION', 3),
('ACHATS / APPROVISIONNEMENT', 3),
('LOGISTIQUE', 3),
('VENTES', 4),
('CONTROLE DE QUALITE', 4),
('RESSOURCES HUMAINES', 4),
('SERVICE CLIENT', 5);



INSERT INTO Posts (postName, pPriority, departementId) VALUES 
('PDG', 1, 1),
('DGA', 1, 1),
('Responsable Stratégie', 2, 1),

('Chef Comptable', 3, 2),
('Comptable Fournisseurs/Clients', 4, 2),
('Assistant Comptable', 5, 2),
('Auditeur Interne', 6, 2),

('Directeur de la Recherche et Développement', 3, 3),
('Parfumeur (Nez)', 4, 3),
('Chimiste Formulateur', 5, 3),
('Technicien de Laboratoire', 6, 3),

('Directeur Marketing', 3, 4),
('Community Manager', 4, 4),

('Directeur Commercial', 3, 5),
('Responsable des Ventes', 4, 5),

('Directeur de Production', 3, 6),
('Responsable de la Chaîne de Production', 4, 6),

('Responsable Qualité', 3, 7),
('Responsable des Conformités et Réglementations', 4, 7),

('Responsable Achats', 3, 8),
('Gestionnaire des Stocks', 4, 8),

('Directeur Logistique', 3, 9),
('Chef de Parc', 4, 9),
('Transporteur / Livreur', 5, 9),
('Planificateur Logistique', 6, 9),

('Directeur des Ressources Humaines', 3, 10),
('Chargé de Recrutement', 4, 10),
('Assistant RH', 5, 10),
('Gestionnaire de Paie', 6, 10),

('Responsable du Service Client', 3, 11),
('Chargé de Relations Clients', 4, 11),
('Agent de Service Après-Vente (SAV)', 5, 11),
('Gestionnaire de Réclamations', 6, 11);





INSERT INTO Employees (firstName, lastName, mail, address, phone, hireDate, salary, departementId, postId) VALUES
('Romain', 'Jollivet', 'romain.jollivet@company.com', '1 Rue de la République, Paris', '0601234567', '2022-01-01', '3747', 1, 1),
('Godwin', 'Oblasse', 'godwin.oblasse@company.com', '2 Avenue des Champs-Élysées, Paris', '0602345678', '2022-02-01', '3747', 1, 2),
('Paul', 'Martin', 'paul.martin@company.com', '3 Boulevard Saint-Germain, Paris', '0603456789', '2022-03-01', '3347', 1, 3),

('Claire', 'Bernard', 'claire.bernard@company.com', '4 Rue de Rivoli, Paris', '0604567890', '2022-04-01', '2947', 2, 4),
('Julien', 'Moreau', 'julien.moreau@company.com', '5 Place Vendôme, Paris', '0605678901', '2022-05-01', '2547', 2, 5),
('Sophie', 'Dubois', 'sophie.dubois@company.com', '6 Quai de la Seine, Paris', '0606789012', '2022-06-01', '2147', 2, 6),
('Emilie', 'Laurent', 'emilie.laurent@company.com', '7 Rue Lafayette, Paris', '0607890123', '2022-07-01', '1747', 2, 7),

('Lucas', 'Perrin', 'lucas.perrin@company.com', '8 Rue Montmartre, Paris', '0608901234', '2022-08-01', '2947', 3, 8),
('Elodie', 'Roux', 'elodie.roux@company.com', '9 Rue Oberkampf, Paris', '0609012345', '2022-09-01', '2547', 3, 9),
('Antoine', 'Fournier', 'antoine.fournier@company.com', '10 Rue des Abbesses, Paris', '0610123456', '2022-10-01', '2147', 3, 10),
('Isabelle', 'Girard', 'isabelle.girard@company.com', '11 Rue de Clichy, Paris', '0611234567', '2022-11-01', '1747', 3, 11),

('Céline', 'Lemoine', 'celine.lemoine@company.com', '12 Rue Saint-Honoré, Paris', '0612345678', '2022-12-01', '2947', 4, 12),
('Thomas', 'Blanc', 'thomas.blanc@company.com', '13 Rue de la Paix, Paris', '0613456789', '2023-01-01', '2547', 4, 13),

('Nathalie', 'Marchand', 'nathalie.marchand@company.com', '14 Rue de lUniversité, Paris', '0614567890', '2023-02-01', '2947', 5, 14),
('François', 'Gautier', 'francois.gautier@company.com', '15 Rue Mouffetard, Paris', '0615678901', '2023-03-01', '2547', 5, 15),

('Camille', 'Henry', 'camille.henry@company.com', '16 Rue de Belleville, Paris', '0616789012', '2023-04-01', '2947', 6, 16),
('Vincent', 'Dupuy', 'vincent.dupuy@company.com', '17 Rue Aubervilliers, Paris', '0617890123', '2023-05-01', '2547', 6, 17),

('Amélie', 'Bazin', 'amelie.bazin@company.com', '18 Rue Saint-Martin, Paris', '0618901234', '2023-06-01', '2947', 7, 18),
('Laurent', 'Michel', 'laurent.michel@company.com', '19 Rue Montorgueil, Paris', '0619012345', '2023-07-01', '2547', 7, 19),

('Julie', 'Renaud', 'julie.renaud@company.com', '20 Rue du Faubourg Saint-Honoré, Paris', '0620123456', '2023-08-01', '2947', 8, 20),
('Pierre', 'Carpentier', 'pierre.carpentier@company.com', '21 Rue de la Convention, Paris', '0621234567', '2023-09-01', '2547', 8, 21),

('Sandrine', 'Lambert', 'sandrine.lambert@company.com', '22 Rue du Louvre, Paris', '0622345678', '2023-10-01', '2947', 9, 22),
('Mickael', 'Bouchard', 'mickael.bouchard@company.com', '23 Rue des Pyrénées, Paris', '0623456789', '2023-11-01', '2547', 9, 23),
('Aurélie', 'Robin', 'aurelie.robin@company.com', '24 Rue Saint-Denis, Paris', '0624567890', '2023-12-01', '2147', 9, 24),
('Guillaume', 'Masson', 'guillaume.masson@company.com', '25 Rue Lecourbe, Paris', '0625678901', '2024-01-01', '1747', 9, 25),

('Nicolas', 'Gerard', 'nicolas.gerard@company.com', '26 Rue de la Roquette, Paris', '0626789012', '2024-02-01', '2947', 10, 26),
('Claire', 'Fabre', 'claire.fabre@company.com', '27 Rue de la Bastille, Paris', '0627890123', '2024-03-01', '2547', 10, 27),
('Benoit', 'Joly', 'benoit.joly@company.com', '28 Rue Monge, Paris', '0628901234', '2024-04-01', '2147', 10, 28),
('Amandine', 'Guerin', 'amandine.guerin@company.com', '29 Rue des Archives, Paris', '0629012345', '2024-05-01', '1747', 10, 29),

('Alexandre', 'Noel', 'alexandre.noel@company.com', '30 Rue des Rosiers, Paris', '0630123456', '2024-06-01', '2947', 11, 30),
('Valérie', 'Rossi', 'valerie.rossi@company.com', '31 Rue de Bercy, Paris', '0631234567', '2024-07-01', '2547', 11, 31),
('Adrien', 'Moulin', 'adrien.moulin@company.com', '32 Rue de la Victoire, Paris', '0632345678', '2024-08-01', '2147', 11, 32),
('Emmanuel', 'Yohore', 'emmanuel.yohore@company.com', '33 Rue de Charenton, Paris', '0633456789', '2024-09-01', '755', 11, 33);
