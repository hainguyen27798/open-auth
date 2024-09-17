package query

const InsertNewToken = `
	INSERT INTO tokens (id, user_id, session, refresh_token)
	VALUES (UUID(), :user_id, :session, :refresh_token)
`

const GetTokenBySession = `SELECT * FROM tokens WHERE session = ? LIMIT 1`

const UpdateRefreshToken = `UPDATE tokens SET refresh_token = ? WHERE id = ?`

const RemoveToken = `DELETE FROM tokens WHERE refresh_token = ?`

const CheckOldRefreshTokenExists = `
	SELECT EXISTS(
		SELECT 1 FROM refresh_tokens_used WHERE refresh_token = ?
	)
`

const CacheOldRefreshToken = `
	INSERT INTO refresh_tokens_used (id, token_id, refresh_token) VALUES (UUID(), ?, ?)
`
