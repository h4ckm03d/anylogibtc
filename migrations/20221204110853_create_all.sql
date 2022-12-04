-- create "transactions" table
CREATE TABLE "transactions" ("id" serial NOT NULL, "datetime" timestamptz NOT NULL, "amount" numeric(6,2) NOT NULL, PRIMARY KEY ("id", "datetime"));
-- create index "transactions_datetime_idx" to table: "transactions"
CREATE INDEX "transactions_datetime_idx" ON "transactions" ("datetime" DESC);


SELECT
  *
FROM
  create_hypertable(
    'transactions',
    'datetime',
    chunk_time_interval = > INTERVAL '1 hour'
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
