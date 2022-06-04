-- name: CreateService :execresult
INSERT INTO services (
   shortcode_id, user_id, service_name, service_id,service_interface, service, service_type, product_id, node_id, subscription_id, subscription_description, base_url, datasync_endpoint, notification_endpoint, network_type
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?, ?, ?
);

-- name: GetService :one
SELECT * FROM services
WHERE id = ? LIMIT 1;

-- name: GetServiceByUserId :one
SELECT * FROM services
WHERE user_id = ? && service=? LIMIT 1;


-- name: GetServiceByShortcodeId :one
SELECT * FROM services
WHERE shortcode_id = ? && service= ? && notification_endpoint IS NOT NULL LIMIT 1;

-- name: ListService :many
SELECT * FROM services
ORDER BY id
LIMIT ?
OFFSET ?;


-- name: UpdateService :execresult
UPDATE services SET  service_name= ?, base_url= ?, datasync_endpoint = ?, notification_endpoint = ?
WHERE id = ?;

-- name: UpdateNotifyEndpointById :execresult
UPDATE services SET notification_endpoint = ?
WHERE id = ?;

-- name: DeleteService :exec
DELETE FROM services
WHERE id = ?;