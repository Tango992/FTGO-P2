CREATE DATABASE GameStore;

CREATE TABLE games (
    game_id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    genre VARCHAR(255) NOT NULL,
    price FLOAT NOT NULL,
    stock INT NOT NULL
);

CREATE TABLE branches (
    branch_id int AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    location VARCHAR(255) NOT NULL
);

CREATE TABLE sales (
    sale_id INT AUTO_INCREMENT PRIMARY KEY,
    game_id INT,
    branch_id INT,
    sale_date DATE NOT NULL,
    quantity INT NOT NULL,
    FOREIGN KEY (game_id) REFERENCES games(game_id),
    FOREIGN KEY (branch_id) REFERENCES branches(branch_id)
);

INSERT INTO games (title, genre, price, stock) VALUES
('Final Fantasy', 'RPG', 59.99, 100),
('FIFA 2023', 'Sports', 49.99, 120),
('Doom Eternal', 'FPS', 49.99, 80);

INSERT INTO branches (name, location) VALUES
('Downtown Branch', '123 Downtown St'),
('Uptown Branch', '456 Uptown Ave');

INSERT INTO sales (game_id, branch_id, sale_date, quantity) VALUES
(1, 1, '2023-08-26', 2),
(2, 2, '2023-08-25', 3);