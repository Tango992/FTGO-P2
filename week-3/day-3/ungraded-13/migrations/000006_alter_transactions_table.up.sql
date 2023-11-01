ALTER TABLE transactions
    ADD COLUMN IF NOT EXISTS store_id INT NOT NULL;

ALTER TABLE transactions
    ADD CONSTRAINT fk_store_id FOREIGN KEY (store_id) REFERENCES stores(id);