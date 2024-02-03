CREATE TABLE IF NOT EXISTS movies (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    added_at TIMESTAMP,
    digitized_at TIMESTAMP
);