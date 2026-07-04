-- +goose Up
CREATE TABLE IF NOT EXISTS transactions_categories (
  transaction_id UUID REFERENCES transactions(transaction_id) ON DELETE CASCADE,
  category_id UUID REFERENCES categories(category_id) ON DELETE CASCADE,
  PRIMARY KEY (transaction_id, category_id)
);

--- Index for the secondary lookups (finding all transactions for a category)
CREATE INDEX IF NOT EXISTS idx_trans_cat_category_id ON transactions_categories(category_id);

-- +goose Down
DROP TABLE IF EXISTS transactions_categories;

