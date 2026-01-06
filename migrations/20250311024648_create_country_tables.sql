-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS countries (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    name VARCHAR(100) NOT NULL,
    iso_code CHAR(3) NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO countries (name, iso_code) VALUES
    ('Vietnam', 'VNM'),
    ('China', 'CHN'),
    ('United States', 'USA');
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS country_dial_codes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    country_id INT NOT NULL,
    dial_code VARCHAR(4) NOT NULL UNIQUE,
    FOREIGN KEY (country_id) REFERENCES countries(id) ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO country_dial_codes (country_id, dial_code) VALUES
    (1, '84'),
    (2, '86'),
    (3, '1');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS country_dial_codes;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS countries;
-- +goose StatementEnd
