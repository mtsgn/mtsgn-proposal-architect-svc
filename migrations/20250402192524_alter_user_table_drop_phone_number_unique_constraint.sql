-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
DROP CONSTRAINT phone_number;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
ADD CONSTRAINT phone_number UNIQUE (phone_number);
-- +goose StatementEnd
