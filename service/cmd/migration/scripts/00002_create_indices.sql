-- +goose Up
CREATE INDEX book_published ON book USING btree (year_published);
CREATE INDEX book_rating ON book USING btree (rating);
CREATE INDEX book_pages ON book USING btree (pages);
CREATE INDEX book_genre_id ON book USING btree (genre_id);
CREATE INDEX book_author_id ON book USING btree (author_id);

-- +goose Down
DROP INDEX "book_published";
DROP INDEX "book_rating";
DROP INDEX "book_pages";
DROP INDEX "book_genre_id";
DROP INDEX "book_author_id";