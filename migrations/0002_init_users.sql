-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id      SERIAL PRIMARY KEY,
    name    VARCHAR(255) NOT NULL,
    job     VARCHAR(255) DEFAULT ''
);

INSERT INTO users (name, job) VALUES 
    ('ed', 'backend dev'), 
    ('hacker', 'devops'), 
    ('erikray', 'backend dev'), 
    ('hank', '');
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS items;
-- +goose StatementEnd
