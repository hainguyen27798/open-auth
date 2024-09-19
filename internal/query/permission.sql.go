package query

const GetAllPermissions = `SELECT * FROM permissions`

var SearchPermissionsBy = map[string]string{
	"service_name": `SELECT * FROM permissions WHERE service_name LIKE ? LIMIT ? OFFSET ?`,
	"resource":     `SELECT * FROM permissions WHERE resource LIKE ? LIMIT ? OFFSET ?`,
	"action":       `SELECT * FROM permissions WHERE action LIKE ? LIMIT ? OFFSET ?`,
	"attributes":   `SELECT * FROM permissions WHERE attributes LIKE ? LIMIT ? OFFSET ?`,
}

var CountPermissionSearchBy = map[string]string{
	"service_name": `SELECT COUNT(id) FROM permissions WHERE service_name LIKE ?`,
	"resource":     `SELECT COUNT(id) FROM permissions WHERE resource LIKE ?`,
	"action":       `SELECT COUNT(id) FROM permissions WHERE action LIKE ?`,
	"attributes":   `SELECT COUNT(id) FROM permissions WHERE attributes LIKE ?`,
}

const GetPermissionById = `SELECT * FROM permissions WHERE id = ? LIMIT 1`

const InsertNewPermission = `
	INSERT INTO permissions (id, service_name, resource, action, attributes, description)
	VALUES (UUID(), :serviceName, :resource, :action, :attributes, :description)
`

const DeletePermission = `DELETE FROM permissions WHERE id = ?`

const UpdatePermission = `UPDATE permissions %s WHERE id = :id`

const SelectPermissionOptions = `
	select p.* from permissions as p
	left join (
		select permission_id from roles_permissions where role_id = ?
	) as rp
	on rp.permission_id = p.id
	where rp.permission_id is null
`
