-- +goose Up
CREATE TABLE hash_records (
    id SERIAL PRIMARY KEY,
    status VARCHAR(50),
    hash  INTEGER
);

-- +goose Down
DROP TABLE hash_records;
