-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = ?
LIMIT 1;

-- name: CreateNewUser :exec
INSERT INTO users (id, name, email, password, status, social_provider, image, verification_code)
VALUES (?, ?, ?, ?, ?, ?, ?, ?);
