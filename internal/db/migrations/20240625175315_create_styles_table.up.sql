CREATE TABLE styles (
    id SERIAL PRIMARY KEY,
    author_id INTEGER NOT NULL REFERENCES authors(id),
    name VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    image BYTEA NOT NULL,
    likes INTEGER NOT NULL DEFAULT 0,
    links TEXT NOT NULL,
    creation_date TIMESTAMP NOT NULL DEFAULT NOW(),
    style_type VARCHAR(50) NOT NULL
);
