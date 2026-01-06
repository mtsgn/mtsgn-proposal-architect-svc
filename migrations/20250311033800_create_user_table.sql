-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id CHAR(36) NOT NULL DEFAULT (UUID()) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    last_login_at TIMESTAMP,
    phone_number VARCHAR(16) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    country_id INT NOT NULL,
    country_dial_code_id INT NOT NULL,

    FOREIGN KEY (country_id) REFERENCES countries(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (country_dial_code_id) REFERENCES country_dial_codes(id) ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
