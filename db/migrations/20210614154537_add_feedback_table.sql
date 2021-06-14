-- +goose Up
-- +goose StatementBegin
SET search_path TO reaction;

CREATE SEQUENCE feedback_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

ALTER SEQUENCE feedback_id_seq
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

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE feedback;
DROP SEQUENCE feedback_id_seq;
-- +goose StatementEnd
