ALTER TABLE users
    ADD COLUMN role_id INT NOT NULL;

ALTER TABLE users
    ADD CONSTRAINT fk_role_id FOREIGN KEY (role_id) REFERENCES roles(id);