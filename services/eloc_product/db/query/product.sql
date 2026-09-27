-- name: CreateProduct :one
INSERT INTO products (
    category_id, name, slug, sku, price, stock, description, status, attributes
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING id, category_id, name, slug, sku, price, stock, description, status, attributes, created_at, updated_at;

-- name: GetProductByID :one
SELECT id, category_id, name, slug, sku, price, stock, description, status, attributes, created_at, updated_at
FROM products
WHERE id = $1 LIMIT 1;

-- name: GetProductBySlug :one
SELECT id, category_id, name, slug, sku, price, stock, description, status, attributes, created_at, updated_at
FROM products
WHERE slug = $1 LIMIT 1;

-- name: GetProductByName :one
SELECT id, category_id, name, slug, sku, price, stock, description, status, attributes, created_at, updated_at
FROM products
WHERE name = $1 LIMIT 1;

-- name: ListProducts :many
SELECT * FROM products
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: ListProductsByCategory :many
SELECT id, category_id, name, slug, sku, price, stock, description, status, created_at
FROM products
WHERE category_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: UpdateProduct :one
UPDATE products
SET 
    category_id = $2,
    name = $3,
    slug = $4,
    sku = $5,
    price = $6,
    stock = $7,
    description = $8,
    status = $9,
    attributes = $10,
    updated_at = now()
WHERE id = $1
RETURNING *;

-- name: UpdateProductStock :one
UPDATE products
SET 
    stock = stock + $2,
    updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteProduct :one
DELETE FROM products
WHERE id = $1
RETURNING *;