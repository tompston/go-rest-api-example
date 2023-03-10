-- Code generated by gomarvin, v0.6.x

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


-- name: Balance_GetWhereUserIdEquals :one
SELECT      b.balance
FROM        balances b
JOIN        transactions t ON b.user_id = t.sender_id OR b.user_id = t.receiver_id
WHERE       b.user_id = $1
GROUP BY    b.balance;


-- cc720122-04df-47c3-98ab-854bdedb9f8c
SELECT      SUM(t.amount)
FROM        transactions t
WHERE       t.sender_id = 'cc720122-04df-47c3-98ab-854bdedb9f8c' 
      OR t.receiver_id = 'cc720122-04df-47c3-98ab-854bdedb9f8c'
;

-- cc720122-04df-47c3-98ab-854bdedb9f8c
SELECT      SUM(t.amount)
FROM        transactions t
WHERE       t.sender_id = 'cc720122-04df-47c3-98ab-854bdedb9f8c';


-- name: Balance_UserReceivedTotalAmount :one
SELECT      SUM(t.amount)
FROM        transactions t
WHERE       t.receiver_id = 'cc720122-04df-47c3-98ab-854bdedb9f8c';

-- name: Balance_UserSentTotalAmount :one
SELECT      SUM(t.amount)
FROM        transactions t
WHERE       t.sender_id = 'cc720122-04df-47c3-98ab-854bdedb9f8c';



-- name: Balance_GetUserBalanceByUserID :one
SELECT    SUM(CASE WHEN transactions.sender_id = @user_id::uuid THEN -transactions.amount ELSE transactions.amount END) as balance
FROM      transactions
WHERE     transactions.sender_id = @user_id::uuid OR transactions.receiver_id = @user_id::uuid;


---------------- Placeholder queries
-- Return 2 rows of the results
SELECT      SUM(t.amount)
FROM        transactions t
WHERE       t.receiver_id = 'cc720122-04df-47c3-98ab-854bdedb9f8c'
    UNION 
SELECT      SUM(t.amount)
FROM        transactions t
WHERE       t.sender_id = 'cc720122-04df-47c3-98ab-854bdedb9f8c';

-- Return balance for placeholder user
SELECT      SUM(CASE WHEN transactions.sender_id = 'cc720122-04df-47c3-98ab-854bdedb9f8c' THEN -transactions.amount ELSE transactions.amount END) as balance
FROM        transactions
WHERE       transactions.sender_id = 'cc720122-04df-47c3-98ab-854bdedb9f8c' OR transactions.receiver_id = 'cc720122-04df-47c3-98ab-854bdedb9f8c';