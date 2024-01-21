-- name: ListsBubbleChat :many
SELECT * FROM bubbles_chat
ORDER BY id;

-- name: GetBubbleChatById :one
SELECT * FROM bubbles_chat
WHERE id = ? LIMIT 1;

-- name: GetBubbleChatByKey :one
SELECT * FROM bubbles_chat
WHERE `key` = ? LIMIT 1;