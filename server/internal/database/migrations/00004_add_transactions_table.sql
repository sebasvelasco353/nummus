-- +goose Up
CREATE TABLE IF NOT EXISTS transactions (
  transaction_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  amount INTEGER NOT NULL,
  name TEXT NOT NULL,
  description TEXT,
  store TEXT
);

-- +goose Down
DROP TABLE IF EXISTS transactions;
