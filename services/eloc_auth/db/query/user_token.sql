-- name: CreateUserToken :one
INSERT INTO user_tokens (
  user_id, refresh_token, expires_at
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetUserToken :one
SELECT * FROM user_tokens
WHERE refresh_token = $1 LIMIT 1;

-- name: DeleteUserToken :exec
DELETE FROM user_tokens
WHERE refresh_token = $1;

-- name: DeleteAllUserTokens :exec
DELETE FROM user_tokens
WHERE user_id = $1;

-- name: DeleteExpiredTokens :exec
DELETE FROM user_tokens
WHERE expires_at < now();
