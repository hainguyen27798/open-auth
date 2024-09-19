package query

const InsertNewRole = `INSERT INTO roles (id, name, description) VALUES (UUID(), :name, :description)`

const GetAllRoles = `SELECT * FROM roles where name LIKE ? LIMIT ? OFFSET ?`

const GetAllRolesTotal = `SELECT COUNT(id) FROM roles where name LIKE ?`

const GetRoleById = `SELECT * FROM roles WHERE id = ? LIMIT 1`

const DeleteRole = `DELETE FROM roles WHERE id = ?`

const UpdateRole = `UPDATE roles %s WHERE id = :id`

const InsertRolePermission = `INSERT INTO roles_permissions (role_id, permission_id) VALUES (?, ?)`

const SelectPermissionByRoleId = `
	SELECT p.id, p.service_name, p.resource, p.action, p.attributes, p.created_at, p.updated_at
	FROM permissions AS p, roles_permissions AS rp
	WHERE rp.role_id = ? AND rp.permission_id = p.id
`

const DeleteRolePermission = `DELETE FROM roles_permissions WHERE role_id = ? AND permission_id = ?`
