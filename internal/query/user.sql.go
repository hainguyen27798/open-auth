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

const GetUserByEmailAndScope = `SELECT * FROM users WHERE email = ? AND scope = ? LIMIT 1`

var SearchUserBy = map[string]string{
	"name": `
		SELECT ur.*, r.name AS role 
		FROM users AS ur 
    	LEFT JOIN roles AS r ON ur.role_id = r.id 
		WHERE ur.name LIKE ? and ur.scope = 'user'
		LIMIT ? OFFSET ?
	`,
	"email": `
		SELECT ur.*, r.name AS role
		FROM users AS ur 
		LEFT JOIN roles AS r ON ur.role_id = r.id
		WHERE ur.email LIKE ? and ur.scope = 'user' 
		LIMIT ? OFFSET ?
	`,
	"status": `
		SELECT ur.*, r.name AS role
		FROM users AS ur 
		LEFT JOIN roles AS r ON ur.role_id = r.id
		WHERE ur.status LIKE ? and ur.scope = 'user'
		LIMIT ? OFFSET ?
	`,
	"social_provider": `
		SELECT ur.*, r.name AS role
		FROM users AS ur 
		LEFT JOIN roles AS r ON ur.role_id = r.id
		WHERE ur.social_provider LIKE ? and ur.scope = 'user' 
		LIMIT ? OFFSET ?
	`,
}

var CountSearchUserBy = map[string]string{
	"name":            `SELECT COUNT(id) FROM users WHERE name LIKE ? and scope = 'user'`,
	"email":           `SELECT COUNT(id) FROM users WHERE email LIKE ? and scope = 'user'`,
	"status":          `SELECT COUNT(id) FROM users WHERE name LIKE ? and scope = 'user'`,
	"social_provider": `SELECT COUNT(id) FROM users WHERE social_provider LIKE ? and scope = 'user'`,
}
