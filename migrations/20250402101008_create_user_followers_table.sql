-- +goose Up
CREATE TABLE IF NOT EXISTS user_followers (
    id CHAR(36) NOT NULL DEFAULT (UUID()),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    user_id CHAR(36) NOT NULL,
    following_user_id CHAR(36) NOT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY unique_user_followers (user_id, following_user_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (following_user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS user_followers;