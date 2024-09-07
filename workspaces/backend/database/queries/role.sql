-- name: InsertNewRole :exec
INSERT INTO roles (id, name, description)
VALUES (UUID(), ?, ?);

-- name: GetAllRoles :many
SELECT *
FROM roles;

-- name: DeleteRole :execrows
DELETE
FROM roles
WHERE id = ?;

-- name: UpdateRole :execrows
UPDATE roles
SET description = ?
WHERE id = ?;