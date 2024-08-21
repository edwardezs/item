-- +goose Up
-- +goose StatementBegin
CREATE TABLE table_lock_status (
    table_name VARCHAR PRIMARY KEY,
    read_only BOOLEAN NOT NULL DEFAULT FALSE
);

INSERT INTO table_lock_status (table_name)
SELECT tablename
FROM pg_tables
WHERE schemaname = 'public' AND tablename NOT IN ('table_lock_status', 'goose_db_version');
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS table_lock_status;
-- +goose StatementEnd
