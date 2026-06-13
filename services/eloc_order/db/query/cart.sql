-- name: CreateCart :one
INSERT INTO carts (
  user_id, product_id, quantity
) VALUES (
  $1, $2, $3
)
ON CONFLICT (user_id, product_id)
DO UPDATE SET 
quantity = carts.quantity + EXCLUDED.quantity,
updated_at = now()
RETURNING *;


-- name: GetUserCart :many
SELECT * FROM carts
WHERE user_id = $1
ORDER BY updated_at DESC;

-- name: UpdateCartQuantity :one
UPDATE carts
SET 
quantity = $3,
updated_at = now()
WHERE user_id = $1 AND product_id = $2
RETURNING *;

-- name: RemoveFromCart :exec
DELETE FROM carts
WHERE user_id = $1 AND product_id = $2;

-- name: ClearUserCart :exec
DELETE FROM carts
WHERE user_id = $1;
