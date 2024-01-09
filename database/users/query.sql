-- name: ListUsers :many
SELECT * FROM users
ORDER BY id;

-- name: GetUserByAuthTicket :one
SELECT * FROM users
WHERE auth_ticket = ?;