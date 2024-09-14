package sql

var GetAllPermissionsBy = map[string]string{
	"service_name": `SELECT * FROM permissions WHERE service_name like ?`,
	"resource":     `SELECT * FROM permissions WHERE resource like ?`,
	"action":       `SELECT * FROM permissions WHERE action like ?`,
	"attributes":   `SELECT * FROM permissions WHERE attributes like ?`,
}
