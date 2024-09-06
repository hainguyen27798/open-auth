-- name: CreateRefreshTokenUsed :exec
INSERT INTO refresh_tokens_used (id, token_id, refresh_token)
VALUES (?, ?, ?);
