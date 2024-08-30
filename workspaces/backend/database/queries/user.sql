-- name: GetUserByEmail :one
SELECT id, email FROM users WHERE email = ? LIMIT 1;