-- name: CreateNewToken :exec
INSERT INTO tokens (id, user_id, session, refresh_token)
VALUES (UUID(), ?, ?, ?);

-- name: GetTokenBySession :one
SELECT id, session, refresh_token
FROM tokens
WHERE session = ?
LIMIT 1;

-- name: UpdateRefreshToken :exec
UPDATE tokens
SET refresh_token = ?
WHERE id = ?;
