-- name: GetAllCategories :many
SELECT * FROM categories;

-- name: GetCategoryById :one
SELECT * FROM categories WHERE id=? LIMIT 1;

-- name: CreateCategory :execresult
INSERT INTO categories (id,name,created_at,updated_at)
VALUES(?,?,now(),now());

-- name: UpdateCategory :execresult
UPDATE categories SET name=? WHERE id=?;

-- name: DeleteCategory :execresult
DELETE FROM categories WHERE id=?;
