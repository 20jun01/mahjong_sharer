-- name: GetUser :one
SELECT name FROM user WHERE id = ?;
