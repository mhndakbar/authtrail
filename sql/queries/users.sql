-- name: CreateUser :exec
INSERT INTO users (id, created_at, updated_at, name, api_key, password)
VALUES (
   $1,
   $2,
   $3,
   $4,
   $5,
   $6
);
--

-- name: GetUser :one
SELECT * FROM users WHERE api_key = $1;
--

-- name: GetUserByName :one
SELECT * FROM users WHERE name = $1;
--

-- name: GetUserByNameAndPassword :one
SELECT * FROM users WHERE name = $1 AND password = $2;
--
