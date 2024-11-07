-- +migrate Up
CREATE TABLE shopping_memos (
  id SERIAL PRIMARY KEY NOT NULL,
  user_id UUID NOT NULL,
  category VARCHAR(10),
  description VARCHAR(255),
  picked boolean,
  created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
  UNIQUE(id, user_id)
);

-- +migrate Down
DROP TABLE IF EXISTS shopping_memos;
