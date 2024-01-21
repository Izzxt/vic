-- name: ListUserStats :many
SELECT * FROM users_stats
ORDER BY id;

-- name: GetUserStats :one
SELECT * FROM users_stats
WHERE user_id = ?;

-- name: UpdateBubbleChat :exec
UPDATE users_stats SET bubble_chat_id = ?
WHERE user_id = ?;

-- name: InsertUserStats :execresult
INSERT INTO users_stats (
  user_id
) VALUES (
  ?
);