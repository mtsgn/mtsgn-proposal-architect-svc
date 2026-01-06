-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS interest (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    type VARCHAR(100) NOT NULL,
    name VARCHAR(100) NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_interest (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    user_id CHAR(36) NOT NULL,
    interest_id INT NOT NULL,

    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (interest_id) REFERENCES interest(id) ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO interest (type, name) VALUES
    ('relation_ship', 'Single'),
    ('relation_ship', 'In a relationship'),
    ('relation_ship', 'Married'),
    ('personality', 'Introvert'),
    ('personality', 'Extrovert'),
    ('personality', 'Ambivert'),
    ('food', 'Milk'),
    ('food', 'Coffee'),
    ('food', 'Tea'),
    ('game', 'LOL'),
    ('game', 'PUBG'),
    ('book_movie_music', 'See you again'),
    ('book_movie_music', 'The sound of silence');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_interest;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS interest;
-- +goose StatementEnd
