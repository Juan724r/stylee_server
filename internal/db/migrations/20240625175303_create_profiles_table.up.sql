CREATE TABLE profiles (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    preferred_styles TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
