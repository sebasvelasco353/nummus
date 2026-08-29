-- +goose Up
CREATE TABLE IF NOT EXISTS refresh_tokens (
  refresh_token_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  refresh_token_hash TEXT NOT NULL UNIQUE,
  user_id UUID REFERENCES users(user_id) ON DELETE CASCADE,
  exp_date TIMESTAMP DEFAULT (CURRENT_TIMESTAMP + INTERVAL '30 days'),
  is_used BOOLEAN DEFAULT FALSE
);

-- +goose Down
DROP TABLE refresh_tokens CASCADE;