
-- name: GetUser :one
SELECT id, first_name, last_name, email, password_hash, row_inserted, row_last_updated FROM user
WHERE id = ?;

-- name: GetUserProfile :one
SELECT id, username, first_name, last_name, email, row_inserted FROM user
WHERE id = ?;

-- name: GetUserByEmail :many
SELECT id, password_hash, username, first_name, last_name FROM user
WHERE email = ?;


-- name: GetUserByUsername :many
SELECT id, password_hash, username, first_name, last_name FROM user
WHERE username = ?;


-- name: InsertUser :execresult
INSERT INTO user (
  username,
  first_name,
  last_name,
  email,
  password_hash,
  row_inserted,
  row_last_updated
) VALUES (
  ?,
  ?,
  ?,
  ?,
  ?,
  NOW(),
  NULL
);
SELECT id FROM user WHERE id = LAST_INSERT_ID();
