-- name: CreateShortCode :execresult
INSERT INTO short_codes (
  short_code, user_id
) VALUES (
  ?,?
);

-- name: GetShortCode :one
SELECT * FROM short_codes
WHERE id = ? LIMIT 1;

-- --name: GetByShortcode :one
-- SELECT id, short_code FROM short_codes
-- WHERE short_code = ? LIMIT 1;

-- name: GetShortcodeByShortCode :one
SELECT id FROM short_codes
WHERE short_code = ? LIMIT 1;

-- name: ListShortCodes :many
SELECT * FROM short_codes
ORDER BY id
LIMIT ?
OFFSET ?;



-- name: UpdateShortCodeName :execresult
UPDATE short_codes SET short_code = ?
WHERE id = ?;

-- name: DeleteShortCode :exec
DELETE FROM short_codes
WHERE id = ?;