CREATE TABLE IF NOT EXISTS stores (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    rating FLOAT NOT NULL CHECK (rating >= 0 AND rating <= 5),
    address VARCHAR(255) NOT NULL,
    lat NUMERIC NOT NULL,
    long NUMERIC NOT NULL
)