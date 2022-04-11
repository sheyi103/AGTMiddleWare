-- name: CreateRole :execresult
INSERT INTO roles (
  name
) VALUES (
  ?
);

-- name: GetRole :one
SELECT * FROM roles
WHERE id = ? LIMIT 1;

-- name: ListRoles :many
SELECT * FROM roles
ORDER BY id
LIMIT ?
OFFSET ?;


-- name: UpdateRoleNames :execresult
UPDATE roles SET name = ?
WHERE id = ?;

-- name: DeleteRole :exec
DELETE FROM roles
WHERE id = ?;