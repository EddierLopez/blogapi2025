-- name: GetUserByEmail :one
SELECT * FROM users WHERE email=? LIMIT 1;

-- name: CreateUser :execresult
INSERT INTO users(name,last_name,email,password,role,created_at,updated_at)
VALUES(?,?,?,?,?,now(),now());