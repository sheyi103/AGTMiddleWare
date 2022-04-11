-- name: CreateUser :execresult
INSERT INTO users (
  name, email, phone_number, contact_person, role_id
) VALUES (
  ?, ?, ?, ?, ?
);

-- name: GetUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT ?
OFFSET ?;


-- name: UpdateUserNames :execresult
UPDATE users SET name = ? , contact_person = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;