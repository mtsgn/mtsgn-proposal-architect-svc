-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
ADD CONSTRAINT unique_phone_constraint UNIQUE (country_dial_code_id, phone_number);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
DROP CONSTRAINT unique_phone_constraint;
-- +goose StatementEnd
