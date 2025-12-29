-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY, 
    
    username TEXT NOT NULL,
    email TEXT NOT NULL, 
    password TEXT NOT NULL, 
    created_at DATETIME NOT NULL DEFAULT (datetime('now')),
    updated_at DATETIME NOT NULL DEFAULT (datetime('now')),

    UNIQUE (id),
    UNIQUE (username),
    UNIQUE (email)
);

CREATE INDEX IF NOT EXISTS idx_credentials_username ON users (username, password);
CREATE INDEX IF NOT EXISTS idx_credentials_email ON users (email, password);
CREATE INDEX IF NOT EXISTS idx_id ON users (id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
