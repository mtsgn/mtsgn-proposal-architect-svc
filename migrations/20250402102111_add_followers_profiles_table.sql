-- +goose Up
ALTER TABLE profile
ADD COLUMN like_count INT NOT NULL DEFAULT 0,
ADD COLUMN following_count INT NOT NULL DEFAULT 0,
ADD COLUMN follower_count INT NOT NULL DEFAULT 0;


-- +goose Down
ALTER TABLE profile
DROP COLUMN like_count,
DROP COLUMN following_count,
DROP COLUMN follower_count;
