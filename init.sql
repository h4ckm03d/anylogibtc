CREATE TABLE
  transactions (
    id SERIAL,
    datetime TIMESTAMPTZ NOT NULL,
    amount numeric(6, 2) NOT NULL,
    PRIMARY KEY(id, datetime)
  );

SELECT
  *
FROM
  create_hypertable(
    'transactions',
    'datetime',
    chunk_time_interval => INTERVAL '1 hour'
  );

CREATE MATERIALIZED VIEW transaction_hourly
WITH
  (timescaledb.continuous) AS
SELECT
  time_bucket(INTERVAL '1 hour', datetime) AS hour,
  SUM(amount) as total
FROM
  transactions
GROUP BY
  hour;