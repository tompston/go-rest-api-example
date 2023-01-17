-- migrate:up
CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


-- user table
CREATE TABLE IF NOT EXISTS users (
  -- init
  user_id       uuid              DEFAULT uuid_generate_v4 () PRIMARY KEY,
  created_at    TIMESTAMP         NOT NULL DEFAULT NOW(),
  updated_at    TIMESTAMP         NOT NULL DEFAULT NOW(),

  -- new columns below
  is_admin     boolean            NOT NULL DEFAULT false,
  username     VARCHAR(300)       NOT NULL UNIQUE,
  email        VARCHAR(300)       NOT NULL UNIQUE,
  password     VARCHAR(700)       NOT NULL
);

-- when the row is updated, update the "updated_at" timestamp
CREATE TRIGGER set_timestamp BEFORE UPDATE ON users
FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();





-- migrate:down

DROP TABLE IF EXISTS users;
