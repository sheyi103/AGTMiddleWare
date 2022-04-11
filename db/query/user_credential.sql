-- name: CreateUserCredential :execresult
INSERT INTO user_credentials (
  client_id, client_secret, shortcode_id, user_id, service_name, service_id, service, service_type, product_id, node_id, subscription_id, subscription_description, base_url, datasyn_endpoint, notification_endpoint, network_type
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
);

-- name: GetUserCredential :one
SELECT * FROM user_credentials
WHERE id = ? LIMIT 1;

-- name: ListUserCredential :many
SELECT * FROM user_credentials
ORDER BY id
LIMIT ?
OFFSET ?;


-- name: UpdateUserCredential :execresult
UPDATE user_credentials SET  service_name= ?, base_url= ?, datasyn_endpoint = ?, notification_endpoint = ?
WHERE id = ?;

-- name: DeleteUserCredential :exec
DELETE FROM user_credentials
WHERE id = ?;