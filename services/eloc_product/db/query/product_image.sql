-- name: AddProductImage :one
INSERT INTO product_images (
    product_id, image_url, is_primary
) VALUES (
    $1, $2, $3
)
RETURNING id, product_id, image_url, is_primary;

-- name: GetProductImages :many
SELECT id, product_id, image_url, is_primary
FROM product_images
WHERE product_id = $1
ORDER BY is_primary DESC, id ASC;

-- name: ResetPrimaryImage :exec
UPDATE product_images
SET is_primary = false
WHERE product_id = $1;

-- name: SetPrimaryImage :one
UPDATE product_images
SET is_primary = true
WHERE id = $1 AND product_id = $2
RETURNING id, product_id, image_url, is_primary;

-- name: DeleteProductImage :exec
DELETE FROM product_images
WHERE id = $1 AND product_id = $2;