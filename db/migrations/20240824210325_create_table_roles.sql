-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS [roles] (
    [id]                    INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL,
    [unique_code]           TEXT UNIQUE NOT NULL,
    [description]           TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS [roles];
-- +goose StatementEnd
