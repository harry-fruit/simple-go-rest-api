-- +goose Up
-- +goose StatementBegin
INSERT INTO [users] (id_status, id_role, [name], [login], [password])
VALUES 
    ((SELECT id FROM [status] WHERE unique_code = 'ACTIVE' LIMIT 1), (SELECT id FROM roles WHERE unique_code = 'ADMIN' LIMIT 1), 'admin', 'admin', 'admin'),
    ((SELECT id FROM [status] WHERE unique_code = 'ACTIVE' LIMIT 1), (SELECT id FROM roles WHERE unique_code = 'STANDARD' LIMIT 1), 'standard', 'standard', 'standard');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM [users] WHERE [login] IN (
    'admin',
    'standard'
);
-- +goose StatementEnd
