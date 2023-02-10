-- name: CreateCommand :one
INSERT INTO commands (
  cmd
) VALUES (
  $1
) RETURNING id;

-- name: GetCommand :one
SELECT * FROM commands
WHERE id = $1 LIMIT 1;

-- name: ListCommands :many
SELECT * FROM commands
ORDER BY id;

-- name: UpdateCommand :one
UPDATE commands
SET cmd = $2
WHERE id = $1
RETURNING *;


-- name: DeleteCommand :exec
DELETE FROM commands
WHERE id = $1;