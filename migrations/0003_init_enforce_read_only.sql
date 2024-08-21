-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION enforce_read_only() RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT read_only FROM table_lock_status WHERE table_name = TG_TABLE_NAME) THEN
        RAISE EXCEPTION 'Table % is read-only', TG_TABLE_NAME;
    END IF;
    RETURN NULL;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP FUNCTION IF EXISTS enforce_read_only();
-- +goose StatementEnd