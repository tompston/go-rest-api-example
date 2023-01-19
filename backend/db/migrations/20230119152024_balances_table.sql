-- migrate:up
CREATE TABLE IF NOT EXISTS balances (
  -- init
  balance_id     uuid             DEFAULT uuid_generate_v4 () PRIMARY KEY,
  created_at    TIMESTAMP         NOT NULL DEFAULT NOW(),
  updated_at    TIMESTAMP         NOT NULL DEFAULT NOW(),

  -- new columns below
  user_id       uuid              NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
  balance       integer           NOT NULL
);

-- when the row is updated, update the "updated_at" timestamp
CREATE TRIGGER set_timestamp BEFORE UPDATE ON balances
FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

-- migrate:down

DROP TABLE IF EXISTS balances;