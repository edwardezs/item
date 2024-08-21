-- +goose Up 
CREATE TABLE table_lock_status (
    table_name VARCHAR PRIMARY KEY,
    read_only BOOLEAN NOT NULL DEFAULT FALSE
);

INSERT INTO table_lock_status (table_name)
VALUES ('items');

-- +goose Down
DROP TABLE table_lock_status