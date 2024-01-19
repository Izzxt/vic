-- name: ListRooms :many
SELECT * FROM rooms;

-- name: GetRoomById :one
SELECT * FROM rooms
WHERE id = ? LIMIT 1;

-- name: GetRoomsByOwnerId :many
SELECT sqlc.embed(r), sqlc.embed(u)
FROM rooms as r
JOIN users as u ON u.id = r.owner_id
WHERE u.id = ?;

-- name: GetActiveRooms :many
SELECT sqlc.embed(r), sqlc.embed(u)
FROM rooms as r
JOIN users as u ON u.id = r.owner_id
WHERE r.users > 0 ORDER BY r.users DESC;

-- name: CreateRoom :execlastid
INSERT INTO rooms (
  owner_id,
  name,
  description,
  max_users,
  model_id,
  flat_category_id,
  trade_mode
) VALUES (?, ?, ?, ?, ?, ?, ?);