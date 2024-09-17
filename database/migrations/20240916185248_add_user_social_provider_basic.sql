-- +goose Up
ALTER TABLE users
    MODIFY COLUMN social_provider enum ('google', 'linkedin', 'basic') NULL DEFAULT 'basic';

UPDATE users
SET social_provider = 'basic'
WHERE social_provider IS NULL;

-- +goose Down
UPDATE users
SET social_provider = NULL
WHERE social_provider = 'basic';

ALTER TABLE users
    MODIFY COLUMN social_provider enum ('google', 'linkedin') NULL;
