-- name: InsertNewRole :exec
INSERT INTO roles (id, name, description)
VALUES (UUID(), ?, ?);

-- name: GetAllRoles :many
SELECT *
FROM roles;

-- name: GetRoleById :one
SELECT *
FROM roles
WHERE id = ?
LIMIT 1;

-- name: DeleteRole :execrows
DELETE
FROM roles
WHERE id = ?;

-- name: UpdateRole :exec
UPDATE roles
SET description = COALESCE(?, description)
WHERE id = ?;