-- name: CreateNewToken :exec
INSERT INTO tokens (id, user_id, session, refresh_token)
VALUES (?, ?, ?, ?);
