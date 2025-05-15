-- name: getPostById :one
SELECT * FROM posts WHERE id=? LIMIT 1;

-- name: getAllPost :many
SELECT * FROM posts;

-- name: createPost :execresult
INSERT INTO posts(id,user_id,category_id,title,content,image,created_at,updated_at)
VALUES(?,?,?,?,?,?,now(),now());