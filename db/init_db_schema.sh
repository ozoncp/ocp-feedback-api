#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
CREATE ROLE service WITH PASSWORD 'service' NOLOGIN;
CREATE SCHEMA reaction;
ALTER SCHEMA reaction OWNER TO postgres;
SET search_path TO reaction;

CREATE SEQUENCE feedback_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

ALTER SEQUENCE feedback_id_seq
    OWNER TO postgres;
    
CREATE SEQUENCE proposal_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

ALTER SEQUENCE proposal_id_seq
    OWNER TO postgres;
    
    
    
CREATE TABLE IF NOT EXISTS feedback
(
    id integer NOT NULL,
    user_id integer NOT NULL,
    classroom_id integer NOT NULL,
    comment text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT feedback_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE feedback
    OWNER to postgres;


CREATE TABLE IF NOT EXISTS proposal
(
    id integer NOT NULL,
    user_id integer NOT NULL,
    lesson_id integer NOT NULL,
    document_id integer NOT NULL,
    CONSTRAINT proposal_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE proposal
    OWNER to postgres;


GRANT CONNECT ON DATABASE postgres TO service;
GRANT USAGE ON SCHEMA reaction TO service;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA reaction TO service;
GRANT SELECT, INSERT, UPDATE, DELETE ON feedback TO service;
GRANT SELECT, INSERT, UPDATE, DELETE ON proposal TO service;
EOSQL