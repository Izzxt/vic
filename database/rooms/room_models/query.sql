-- name: ListRoomModels :many
SELECT * FROM room_models;

-- name: GetModelById :one
SELECT * FROM room_models
WHERE id = ?;