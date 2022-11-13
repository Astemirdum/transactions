CREATE SCHEMA tv1;

CREATE TABLE IF NOT EXISTS tv1.balances (
    id  INT GENERATED ALWAYS AS IDENTITY  NOT NULL,
    user_id   INT            NOT NULL,
    cash            BIGINT                     NOT NULL DEFAULT 0,
    last_transaction TIMESTAMP      NOT NULL DEFAULT (now() at time zone 'utc')
);

