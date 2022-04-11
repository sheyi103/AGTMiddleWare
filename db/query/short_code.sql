-- name: CreateShortCode :execresult
INSERT INTO short_codes (
  short_code
) VALUES (
  ?
);

-- name: GetShortCode :one
SELECT * FROM short_codes
WHERE id = ? LIMIT 1;

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