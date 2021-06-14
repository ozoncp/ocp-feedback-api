-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA reaction;
ALTER SCHEMA reaction OWNER TO postgres;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP SCHEMA reaction;
-- +goose StatementEnd
