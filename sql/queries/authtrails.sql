-- name: CreateAuthTrail :exec
INSERT INTO authtrails (
  id,
  created_at,
  updated_at,
  type,
  user_id
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5
);
--

-- name: GetAuthTrailsForUser :many
SELECT * FROM authtrails WHERE user_id = $1;
--