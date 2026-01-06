-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS profile (
    id CHAR(36) NOT NULL DEFAULT (UUID()),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    user_id CHAR(36) NOT NULL,
    nickname VARCHAR(100),
    gender VARCHAR(100),
    birthday DATE,
    avatar VARCHAR(100),
    voice VARCHAR(100),
    self_intro VARCHAR(100),

    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS profile;
-- +goose StatementEnd
