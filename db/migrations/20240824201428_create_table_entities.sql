-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS entities (
			[id] INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL,
			[unique_code] TEXT UNIQUE NOT NULL,
			[description] TEXT
			);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS entities;
-- +goose StatementEnd
