-- name: CreateProduct :one
INSERT INTO products (
    category_id, name, slug, sku, price, stock, attributes
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
RETURNING id, category_id, name, slug, sku, price, stock, attributes, created_at, updated_at;

-- name: GetProductByID :one
SELECT id, category_id, name, slug, sku, price, stock, attributes, created_at, updated_at
FROM products
WHERE id = $1 LIMIT 1;

-- name: GetProductBySlug :one
SELECT id, category_id, name, slug, sku, price, stock, attributes, created_at, updated_at
FROM products
WHERE slug = $1 LIMIT 1;

-- name: ListProducts :many
SELECT id, category_id, name, slug, sku, price, stock, created_at
FROM products
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: ListProductsByCategory :many
SELECT id, category_id, name, slug, sku, price, stock, created_at
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
    attributes = $8,
    updated_at = now()
WHERE id = $1
RETURNING id, category_id, name, slug, sku, price, stock, attributes, updated_at;

-- name: UpdateProductStock :one
UPDATE products
SET 
    stock = stock + $2,
    updated_at = now()
WHERE id = $1
RETURNING id, name, stock;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;