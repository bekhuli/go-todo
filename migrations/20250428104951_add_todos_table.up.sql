CREATE TABLE IF NOT EXISTS todos (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    title TEXT NOT NULL,
    done BOOLEAN NOT NULL DEFAULT false,

    FOREIGN KEY (user_id) REFERENCES users(id)
);