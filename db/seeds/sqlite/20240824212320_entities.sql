-- +goose Up
-- +goose StatementBegin
INSERT INTO entities ([unique_code], [description]) 
VALUES 
    ('USER', 'user entity');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM entities WHERE [unique_code] IN (
    'USER'
);
-- +goose StatementEnd
