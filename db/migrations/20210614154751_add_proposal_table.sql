-- +goose Up
-- +goose StatementBegin
SET search_path TO reaction;

CREATE SEQUENCE proposal_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

ALTER SEQUENCE proposal_id_seq
    OWNER TO postgres;

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
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE proposal;
DROP SEQUENCE proposal_id_seq;
-- +goose StatementEnd
