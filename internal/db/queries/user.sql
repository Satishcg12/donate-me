-- name: CreateUser :one
INSERT INTO users
(full_name, email, password)
VALUES
(?, ?, ?)
RETURNING *;


-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = ?;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = ?;

-- name: UpdateUser :one
UPDATE users
SET full_name = ?, email = ?, password = ?
WHERE id = ?
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;

-- name: ListUsers :many
SELECT * FROM users ORDER BY id ASC LIMIT ? OFFSET ?;

-- name: CountUsers :one
SELECT COUNT(*) FROM users;