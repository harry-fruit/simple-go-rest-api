-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS [users] (
	[id]                    INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL,
	[id_status]             INTEGER NOT NULL,
	[id_role]               INTEGER NOT NULL,
    [name]                  TEXT NOT NULL,
	[login]                 TEXT UNIQUE NOT NULL,
	[password]              TEXT,
    FOREIGN KEY ([id_status]) REFERENCES [status] (id),
    FOREIGN KEY ([id_role])   REFERENCES [roles] (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS [users];
-- +goose StatementEnd
