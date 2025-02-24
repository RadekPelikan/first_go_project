-- name: GetAllUsers :many
SELECT * FROM user;

-- name: InsertUser :exec
INSERT INTO user (name, email, password)
VALUES (?, ?, ?);
