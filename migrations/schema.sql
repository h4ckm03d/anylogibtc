--
-- PostgreSQL database dump
--

-- Dumped from database version 13.8
-- Dumped by pg_dump version 14.4 (Ubuntu 14.4-0ubuntu0.22.04.1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: timescaledb; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS timescaledb WITH SCHEMA public;


--
-- Name: EXTENSION timescaledb; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION timescaledb IS 'Enables scalable inserts and complex queries for time-series data';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: transactions; Type: TABLE; Schema: public; Owner: timescaledb
--

CREATE TABLE public.transactions (
    id integer NOT NULL,
    datetime timestamp with time zone NOT NULL,
    amount numeric(6,2) NOT NULL
);


ALTER TABLE public.transactions OWNER TO timescaledb;

--
-- Name: _direct_view_3; Type: VIEW; Schema: _timescaledb_internal; Owner: timescaledb
--

CREATE VIEW _timescaledb_internal._direct_view_3 AS
 SELECT public.time_bucket('01:00:00'::interval, transactions.datetime) AS hour,
    sum(transactions.amount) AS total
   FROM public.transactions
  GROUP BY (public.time_bucket('01:00:00'::interval, transactions.datetime));


ALTER TABLE _timescaledb_internal._direct_view_3 OWNER TO timescaledb;

--
-- Name: _hyper_2_1_chunk; Type: TABLE; Schema: _timescaledb_internal; Owner: timescaledb
--

CREATE TABLE _timescaledb_internal._hyper_2_1_chunk (
    CONSTRAINT constraint_1 CHECK (((datetime >= '2019-10-05 07:00:00+00'::timestamp with time zone) AND (datetime < '2019-10-05 08:00:00+00'::timestamp with time zone)))
)
INHERITS (public.transactions);


ALTER TABLE _timescaledb_internal._hyper_2_1_chunk OWNER TO timescaledb;

--
-- Name: _materialized_hypertable_3; Type: TABLE; Schema: _timescaledb_internal; Owner: timescaledb
--

CREATE TABLE _timescaledb_internal._materialized_hypertable_3 (
    hour timestamp with time zone NOT NULL,
    total numeric
);


ALTER TABLE _timescaledb_internal._materialized_hypertable_3 OWNER TO timescaledb;

--
-- Name: _partial_view_3; Type: VIEW; Schema: _timescaledb_internal; Owner: timescaledb
--

CREATE VIEW _timescaledb_internal._partial_view_3 AS
 SELECT public.time_bucket('01:00:00'::interval, transactions.datetime) AS hour,
    sum(transactions.amount) AS total
   FROM public.transactions
  GROUP BY (public.time_bucket('01:00:00'::interval, transactions.datetime));


ALTER TABLE _timescaledb_internal._partial_view_3 OWNER TO timescaledb;

--
-- Name: transaction_hourly; Type: VIEW; Schema: public; Owner: timescaledb
--

CREATE VIEW public.transaction_hourly AS
 SELECT _materialized_hypertable_3.hour,
    _materialized_hypertable_3.total
   FROM _timescaledb_internal._materialized_hypertable_3
  WHERE (_materialized_hypertable_3.hour < COALESCE(_timescaledb_internal.to_timestamp(_timescaledb_internal.cagg_watermark(3)), '-infinity'::timestamp with time zone))
UNION ALL
 SELECT public.time_bucket('01:00:00'::interval, transactions.datetime) AS hour,
    sum(transactions.amount) AS total
   FROM public.transactions
  WHERE (transactions.datetime >= COALESCE(_timescaledb_internal.to_timestamp(_timescaledb_internal.cagg_watermark(3)), '-infinity'::timestamp with time zone))
  GROUP BY (public.time_bucket('01:00:00'::interval, transactions.datetime));


ALTER TABLE public.transaction_hourly OWNER TO timescaledb;

--
-- Name: transactions_id_seq; Type: SEQUENCE; Schema: public; Owner: timescaledb
--

CREATE SEQUENCE public.transactions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.transactions_id_seq OWNER TO timescaledb;

--
-- Name: transactions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: timescaledb
--

ALTER SEQUENCE public.transactions_id_seq OWNED BY public.transactions.id;


--
-- Name: _hyper_2_1_chunk id; Type: DEFAULT; Schema: _timescaledb_internal; Owner: timescaledb
--

ALTER TABLE ONLY _timescaledb_internal._hyper_2_1_chunk ALTER COLUMN id SET DEFAULT nextval('public.transactions_id_seq'::regclass);


--
-- Name: transactions id; Type: DEFAULT; Schema: public; Owner: timescaledb
--

ALTER TABLE ONLY public.transactions ALTER COLUMN id SET DEFAULT nextval('public.transactions_id_seq'::regclass);


--
-- Name: _hyper_2_1_chunk 1_1_transactions_pkey; Type: CONSTRAINT; Schema: _timescaledb_internal; Owner: timescaledb
--

ALTER TABLE ONLY _timescaledb_internal._hyper_2_1_chunk
    ADD CONSTRAINT "1_1_transactions_pkey" PRIMARY KEY (id, datetime);


--
-- Name: transactions transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: timescaledb
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (id, datetime);


--
-- Name: _hyper_2_1_chunk_transactions_datetime_idx; Type: INDEX; Schema: _timescaledb_internal; Owner: timescaledb
--

CREATE INDEX _hyper_2_1_chunk_transactions_datetime_idx ON _timescaledb_internal._hyper_2_1_chunk USING btree (datetime DESC);


--
-- Name: _materialized_hypertable_3_hour_idx; Type: INDEX; Schema: _timescaledb_internal; Owner: timescaledb
--

CREATE INDEX _materialized_hypertable_3_hour_idx ON _timescaledb_internal._materialized_hypertable_3 USING btree (hour DESC);


--
-- Name: transactions_datetime_idx; Type: INDEX; Schema: public; Owner: timescaledb
--

CREATE INDEX transactions_datetime_idx ON public.transactions USING btree (datetime DESC);


--
-- Name: _hyper_2_1_chunk ts_cagg_invalidation_trigger; Type: TRIGGER; Schema: _timescaledb_internal; Owner: timescaledb
--

CREATE TRIGGER ts_cagg_invalidation_trigger AFTER INSERT OR DELETE OR UPDATE ON _timescaledb_internal._hyper_2_1_chunk FOR EACH ROW EXECUTE FUNCTION _timescaledb_internal.continuous_agg_invalidation_trigger('2');


--
-- Name: _materialized_hypertable_3 ts_insert_blocker; Type: TRIGGER; Schema: _timescaledb_internal; Owner: timescaledb
--

CREATE TRIGGER ts_insert_blocker BEFORE INSERT ON _timescaledb_internal._materialized_hypertable_3 FOR EACH ROW EXECUTE FUNCTION _timescaledb_internal.insert_blocker();


--
-- Name: transactions ts_cagg_invalidation_trigger; Type: TRIGGER; Schema: public; Owner: timescaledb
--

CREATE TRIGGER ts_cagg_invalidation_trigger AFTER INSERT OR DELETE OR UPDATE ON public.transactions FOR EACH ROW EXECUTE FUNCTION _timescaledb_internal.continuous_agg_invalidation_trigger('2');


--
-- Name: transactions ts_insert_blocker; Type: TRIGGER; Schema: public; Owner: timescaledb
--

CREATE TRIGGER ts_insert_blocker BEFORE INSERT ON public.transactions FOR EACH ROW EXECUTE FUNCTION _timescaledb_internal.insert_blocker();


--
-- PostgreSQL database dump complete
--

