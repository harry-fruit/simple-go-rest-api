-- +goose Up
-- +goose StatementBegin
INSERT INTO [status] ([id_entity], [unique_code], [description]) 
VALUES 
    ((SELECT id FROM entities WHERE unique_code = 'USER' LIMIT 1), 'ACTIVE', 'user active'),
    ((SELECT id FROM entities WHERE unique_code = 'USER' LIMIT 1), 'INACTIVE', 'user inactive');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM [status] WHERE [unique_code] IN (
    'ACTIVE',
    'INACTIVE'
);
-- +goose StatementEnd
