-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
	[id] INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL,
	[name] TEXT NOT NULL,
	[login] TEXT UNIQUE NOT NULL,
	[password] TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
