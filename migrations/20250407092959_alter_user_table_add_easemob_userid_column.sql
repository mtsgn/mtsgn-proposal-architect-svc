-- +goose Up
-- +goose StatementBegin
-- Make sure you're using your actual database name in place of `your_db_name`
SET @column_exists := (
    SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS
    WHERE table_name = 'users'
    AND table_schema = 'wjfh'
    AND column_name = 'easemob_uuid'
);

-- If column doesn't exist, prepare the ALTER statement
SET @sql := IF(@column_exists = 0,
    'ALTER TABLE users ADD COLUMN easemob_uuid VARCHAR(100) NULL;',
    'SELECT "Column already exists";'
);

-- Run the SQL
PREPARE stmt FROM @sql;
EXECUTE stmt;
DEALLOCATE PREPARE stmt;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
DROP COLUMN easemob_uuid;
-- +goose StatementEnd
