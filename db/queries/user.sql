-- name: GetUserByEmail :one
SELECT * FROM users WHERE email=? LIMIT 1;

-- name: CreateUser :execresult
INSERT INTO users(name,last_name,email,password,role,created_at,updated_at)
VALUES(?,?,?,?,?,now(),now());

-- name: UpdateUser :execresult
UPDATE users SET name=?,last_name=?,updated_at=now() WHERE id=?;

-- name: UpdateUserRole :execresult
UPDATE users set role=? WHERE id=?;

-- name: UpdateUserPassword :execresult
UPDATE users set password=? WHERE id=?;