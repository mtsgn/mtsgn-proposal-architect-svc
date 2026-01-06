-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
ADD COLUMN easemob_username VARCHAR(255) NULL,
ADD COLUMN easemob_password VARCHAR(255) NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
DROP COLUMN easemob_username,
DROP COLUMN easemob_password;
-- +goose StatementEnd
