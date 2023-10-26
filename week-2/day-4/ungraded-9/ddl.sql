CREATE DATABASE p2_ungraded_7;

CREATE TABLE IF NOT EXISTS stores (
    id INT AUTO_INCREMENT NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(100) NOT NULL,
    type VARCHAR(15) DEFAULT "silver",
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS products (
    id INT AUTO_INCREMENT NOT NULL,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(255) NOT NULL,
    image_url VARCHAR(255) NOT NULL,
    price INT NOT NULL,
    store_id INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (store_id) REFERENCES stores(id)
);