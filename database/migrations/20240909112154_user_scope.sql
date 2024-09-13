-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
    ADD scope enum ('user', 'admin') NOT NULL DEFAULT 'user';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
    DROP COLUMN scope;
-- +goose StatementEnd
