-- +goose Up
CREATE INDEX author_full_name_idx ON author USING GIN (to_tsvector('simple', first_name || ' ' || last_name));

-- +goose Down
DROP INDEX author_full_name_idx;