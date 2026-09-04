-- name: CreateUserToken :one
INSERT INTO user_tokens (
  id, user_id, refresh_token,user_agent,client_ip,is_blocked, expires_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: GetUserToken :one
SELECT * FROM user_tokens
WHERE id = $1 LIMIT 1;

-- name: GetUserTokenByRefreshToken :one
SELECT * FROM user_tokens
WHERE refresh_token = $1 LIMIT 1;

-- name: RevokeUserToken :one
UPDATE user_tokens
SET is_blocked = true
WHERE id = $1
RETURNING *;

-- name: DeleteUserToken :exec
DELETE FROM user_tokens
WHERE refresh_token = $1;

-- name: DeleteAllUserTokens :exec
DELETE FROM user_tokens
WHERE user_id = $1;

-- name: DeleteExpiredTokens :exec
DELETE FROM user_tokens
WHERE expires_at < now();
