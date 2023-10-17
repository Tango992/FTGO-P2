CREATE DATABASE p2_ungraded_3;

CREATE TABLE IF NOT EXISTS Status (
    Id INT AUTO_INCREMENT NOT NULL,
    Name VARCHAR(50) NOT NULL,
    PRIMARY KEY (Id)
);

CREATE TABLE IF NOT EXISTS Inventories (
    Id INT AUTO_INCREMENT NOT NULL,
    Name VARCHAR(100) NOT NULL,
    Stock INT UNSIGNED NOT NULL,
    Description VARCHAR(255),
    Status_id INT NOT NULL,
    PRIMARY KEY (Id),
    FOREIGN KEY (Status_id) REFERENCES Status(id)
);

INSERT INTO Status (Name) VALUES ("Active"), ("Broken");

INSERT INTO Inventories (Name, Stock, Description, Status_id) VALUES
("Vibranium", 1, "Exotic material", 1);