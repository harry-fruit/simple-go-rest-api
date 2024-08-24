-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS [status] (
	[id]                    INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL,
    [id_entity]             INTEGER NOT NULL,
	[unique_code]           TEXT UNIQUE NOT NULL,
	[description]           TEXT,
	FOREIGN KEY ([id_entity]) REFERENCES [entities] ([id]) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS [status];
-- +goose StatementEnd
