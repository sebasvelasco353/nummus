-- +goose Up
CREATE TABLE IF NOT EXISTS users (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name TEXT,
  last_name TEXT,
  email TEXT UNIQUE NOT NULL,
  password TEXT
);

-- +goose Down
DROP TABLE users;
