-- name: CreateOrder :one
INSERT INTO orders (
  user_id, total_amount, status, shipping_address, customer_phone
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetOrder :one
SELECT * FROM orders
WHERE id = $1 LIMIT 1;

-- name: ListUserOrders :many
SELECT id, user_id, total_amount, status, shipping_address, customer_phone, created_at FROM orders
WHERE user_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: UpdateOrderStatus :one
UPDATE orders
SET 
    status = $2,
    updated_at = now()
WHERE id = $1
RETURNING *;
