CREATE TABLE users
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users (name, email, password) VALUES
('admin', 'admin@gmail.com', 'admin'),
('user', 'user@gmail.com', 'user');