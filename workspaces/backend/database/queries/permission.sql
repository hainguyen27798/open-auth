-- name: InsertNewPermission :exec
INSERT INTO permissions (id, service_name, resource, action, attributes, description)
VALUES (UUID(), ?, ?, ?, ?, ?)