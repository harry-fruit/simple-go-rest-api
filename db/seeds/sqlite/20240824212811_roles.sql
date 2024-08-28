-- +goose Up
-- +goose StatementBegin
INSERT INTO [roles] ([unique_code], [description])
VALUES 
    ('ADMIN', 'administrator'),
    ('STANDARD', 'standard');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM [roles] WHERE [unique_code] IN (
    'ADMIN',
    'STANDARD'
);
-- +goose StatementEnd
