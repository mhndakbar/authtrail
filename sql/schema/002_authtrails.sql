-- +goose Up
CREATE TABLE authtrails (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  type TEXT NOT NULL,
  user_id UUID NOT NULL,
  CONSTRAINT fk_authtrails_users FOREIGN KEY (user_id) REFERENCES users(id)
);

-- +goose Down
DROP TABLE authtrails;