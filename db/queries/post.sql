-- name: GetPostById :one
SELECT * FROM posts WHERE id=? LIMIT 1;

-- name: GetAllPost :many
SELECT * FROM posts;

-- name: CreatePost :execresult
INSERT INTO posts(id,user_id,category_id,title,content,image,created_at,updated_at)
VALUES(?,?,?,?,?,?,now(),now());

-- name: GetPostsByUser :many
SELECT * FROM posts WHERE user_id=?;

-- name: GetPostsByCategory :many
SELECT * FROM posts WHERE category_id=?;

-- name: DeletePost :execresult
DELETE FROM posts where id=?;