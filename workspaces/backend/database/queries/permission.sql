-- name: InsertNewPermission :exec
INSERT INTO permissions (id, service_name, resource, action, attributes, description)
VALUES (UUID(), ?, ?, ?, ?, ?);

-- name: GetAllPermissions :many
SELECT *
FROM permissions;

-- name: UpdatePermission :exec
UPDATE permissions
SET service_name = COALESCE(sqlc.narg('service_name'), service_name),
    resource     = COALESCE(sqlc.narg('resource'), resource),
    action       = COALESCE(sqlc.narg('action'), action),
    attributes   = COALESCE(sqlc.narg('attributes'), attributes),
    description  = COALESCE(sqlc.narg('description'), description)
WHERE id = ?;

-- name: DeletePermission :execrows
DELETE
FROM permissions
WHERE id = ?;
