-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE IF NOT EXISTS access_tokens (
    id INTEGER PRIMARY KEY  AUTOINCREMENT,

    token TEXT NOT NULL,
    active TINYINT(1) NOT NULL DEFAULT(1),
    created_at DATETIME NOT NULL DEFAULT (datetime('now')),
    updated_at DATETIME NOT NULL DEFAULT (datetime('now')),

    UNIQUE (id),
    UNIQUE (token)
);

CREATE INDEX IF NOT EXISTS idx_at_id ON access_tokens (id);
CREATE INDEX IF NOT EXISTS idx_at_token ON access_tokens (token);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS access_tokens;
-- +goose StatementEnd
