-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = ?
LIMIT 1;

-- name: InsertNewUser :exec
INSERT INTO users (id, name, email, password, status, verification_code, scope)
VALUES (UUID(), ?, ?, ?, 'request', ?, 'user');

-- name: InsertSuperUser :exec
INSERT INTO users (id, name, email, password, status, scope)
VALUES (UUID(), 'Admin', ?, ?, 'active', 'admin');
