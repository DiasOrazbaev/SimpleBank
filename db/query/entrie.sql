-- name: CreateEntrie :one
INSERT INTO entries (
  account_id, amount
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetEntrie :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: ListEntrie :many
SELECT * FROM entries
ORDER BY id
LIMIT $1
OFFSET $2;
