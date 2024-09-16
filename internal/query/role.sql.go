package query

const InsertNewRole = `INSERT INTO roles (id, name, description) VALUES (UUID(), :name, :description)`

const GetAllRoles = `SELECT * FROM roles`

const GetRoleById = `SELECT * FROM roles WHERE id = ? LIMIT 1`

const DeleteRole = `DELETE FROM roles WHERE id = ?`

const UpdateRole = `UPDATE roles %s WHERE id = :id`
