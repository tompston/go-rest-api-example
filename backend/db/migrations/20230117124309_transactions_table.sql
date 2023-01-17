-- migrate:up

CREATE TABLE IF NOT EXISTS transactions (
  -- init
  transaction_id     uuid         DEFAULT uuid_generate_v4 () PRIMARY KEY,
  created_at    TIMESTAMP         NOT NULL DEFAULT NOW(),
  updated_at    TIMESTAMP         NOT NULL DEFAULT NOW(),

  -- new fields
  sender_id     uuid              NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
  receiver_id   uuid              NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
  amount        integer           NOT NULL
);

-- when the row is updated, update the "updated_at" timestamp
CREATE TRIGGER set_timestamp BEFORE UPDATE ON transactions
FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();


-- migrate:down

DROP TABLE IF EXISTS transactions;