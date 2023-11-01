CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    stock BIGINT NOT NULL CHECK (stock >= 0),
    price DECIMAL NOT NULL CHECK (price >= 0)
);