-- name: CacheOldRefreshToken :exec
INSERT INTO refresh_tokens_used (id, token_id, refresh_token)
VALUES (UUID(), ?, ?);

-- name: CheckOldRefreshTokenExists :one
SELECT COUNT(id)
FROM refresh_tokens_used
WHERE refresh_token = ?;
