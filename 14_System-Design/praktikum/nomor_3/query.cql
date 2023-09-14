-- Contoh berkas Cypher untuk Neo4j

-- Data Contoh (pengguna)
CREATE TABLE users (
    user_id UUID PRIMARY KEY,
    username TEXT
);

INSERT INTO users (user_id, username) VALUES (uuid(), 'John');
INSERT INTO users (user_id, username) VALUES (uuid(), 'Alice');
INSERT INTO users (user_id, username) VALUES (uuid(), 'Bob');

-- Ambil data dari pengguna
SELECT * FROM users;