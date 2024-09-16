package query

const InsertBasicUser = `
	INSERT INTO users (id, name, email, password, status, social_provider, verification_code, scope)
	VALUES (UUID(), :name, :email, :password, 'request', 'basic', :verification_code, 'user')
`

const InsertSuperuser = `
	INSERT INTO users (id, name, email, password, status, scope)
	VALUES (UUID(), 'Admin', :email, :password, 'active', 'admin')
`

const CheckUserByEmail = `SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)`

const GetUserByEmail = `SELECT * FROM users WHERE email = ? LIMIT 1`

const GetUserByEmailAndScope = `SELECT * FROM users WHERE email = ? AND scope = ?LIMIT 1`
