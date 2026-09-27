-- +goose Up
CREATE TYPE expense_types AS ENUM ('in', 'out');

CREATE TABLE IF NOT EXISTS transactions (
  transaction_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  account UUID REFERENCES accounts(account_id) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  expense_type expense_types NOT NULL,
  amount BIGINT NOT NULL CHECK (amount > 0),
  name TEXT NOT NULL,
  description TEXT,
  store TEXT
);

--- Index for listing an account's transactions ordered by date
CREATE INDEX IF NOT EXISTS idx_transactions_account_created_at ON transactions(account, created_at);

-- +goose Down
DROP TABLE IF EXISTS transactions;
DROP TYPE IF EXISTS expense_types;
