

-- name: CreateDonation :one
INSERT INTO donations
(full_name, email, message, status, amount)
VALUES
(?, ?, ?, "pending", ?)
RETURNING *;

-- name: GetDonationByID :one
SELECT * FROM donations WHERE id = ?;

-- name: UpdateStatus :one
UPDATE donations
SET status = ?
WHERE id = ?
RETURNING *;

-- name: ListDonations :many
SELECT * FROM donations ORDER BY updated_at DESC LIMIT ? OFFSET ?;

-- name: CountDonations :one
SELECT COUNT(*) FROM donations;

-- name: DeleteDonation :exec
DELETE FROM donations WHERE id = ?;

-- name: UpdateDonation :one
UPDATE donations
SET full_name = ?, email = ?, message = ?, status = ?, amount = ?
WHERE id = ?
RETURNING *;


-- name: GetTotalDonationsAmount :one
SELECT SUM(amount) FROM donations WHERE status = 'COMPLETE';
