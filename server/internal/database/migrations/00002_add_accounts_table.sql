-- +goose Up
CREATE TYPE account_types AS ENUM ('savings', 'debit', 'credit');

CREATE TABLE IF NOT EXISTS accounts (
  account_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  owner UUID REFERENCES users(user_id) ON DELETE CASCADE,
  bank TEXT NOT NULL,
  name TEXT NOT NULL,
  balance INTEGER NOT NULL,
  currency TEXT NOT NULL,
  account_type account_types NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS accounts CASCADE;
DROP TYPE IF EXISTS account_types;

