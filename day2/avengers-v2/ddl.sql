CREATE TABLE heroes (
	id INT PRIMARY KEY AUTO_INCREMENT,
	name VARCHAR(50) NOT NULL,
	universe VARCHAR(100),
	skill VARCHAR(255),
	imageURL VARCHAR(255)
);

CREATE TABLE villain (
	id INT PRIMARY KEY AUTO_INCREMENT,
	name VARCHAR(50),
	universe VARCHAR(100),
	imageURL VARCHAR(255)
);

-- Insert sample data
-- Heroes
INSERT INTO heroes (name, universe, skill, imageURL)
VALUES
	("Hulk", "Mars", "Punch", "hulk.jpg"),
	("Spiderman", "Vanus", "Spider", "spiderman.jpg"),
	("Iron Man", "Yupiter", "Iron", "ironman.jpg"),
	("Thor", "Saturn", "Hammer", "thor.jpg");

-- Villain
INSERT INTO villain (name, universe, imageURL)
VALUES
	("Crocodille", "Mars", "crocodille.jpg"),
	("Comodo", "Vanus", "comodo.jpg"),
	("Lion", "Yupiter", "lion.jpg"),
	("Tiger", "Saturn", "tiger.jpg");
	
-- Create Table Inventories
CREATE TABLE inventories (
	id INT PRIMARY KEY AUTO_INCREMENT,
	name VARCHAR(50),
	stock INT
);

-- Insert sample data to inventories
INSERT INTO inventories (name, stock)
VALUES
	("Axe", 3),
	("Sword", 2),
	("Archer", 5),
	("Gun", 7),
	("Armor", 4);
	
-- Create Table Criminal Reports
CREATE TABLE CriminalReports (
	id INT PRIMARY KEY AUTO_INCREMENT,
	hero_id INT,
	villain_id INT,
	location VARCHAR(255),
	date DATETIME,
	description TEXT,
	status VARCHAR(50),
	
	FOREIGN KEY (hero_id) REFERENCES heroes(id),
	FOREIGN KEY (villain_id) REFERENCES villain(id)
);

-- Insert sample data to Criminal Reports
INSERT INTO CriminalReports (hero_id, villain_id, location, date, description, status)
VALUES
	(1, 3, "Kelapa Dua, Depok", "2023-09-20 13:00:00", "Contoh deskripsi", "success"),
	(2, 2, "Pasar Senen, Jakarta", "2023-08-20 13:00:00", "Contoh deskripsi", "success"),
	(3, 1, "Jatinegara, Jakarta", "2023-10-15 13:00:00", "Contoh deskripsi", "success"),
	(4, 3, "Dago, Bandung", "2023-10-18 13:00:00", "Contoh deskripsi", "on progres");