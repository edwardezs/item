-- +goose Up
CREATE TABLE users
(
    id              SERIAL PRIMARY KEY,
    name            VARCHAR(255) NOT NULL,
    username        VARCHAR(255) NOT NULL,
    password_hash   VARCHAR(255) NOT NULL
);

CREATE TABLE todo_lists
(
    id              SERIAL PRIMARY KEY,
    title           VARCHAR(255) NOT NULL,
    description     VARCHAR(255)
);

CREATE TABLE users_lists
(
    id       SERIAL PRIMARY KEY,
    user_id  INT REFERENCES users (id) ON DELETE CASCADE,
    list_id  INT REFERENCES todo_lists (id) ON DELETE CASCADE
);

CREATE TABLE todo_items
(
    id              SERIAL PRIMARY KEY,
    title           VARCHAR(255) NOT NULL,
    description     VARCHAR(255),
    done            BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE lists_items
(
    id       SERIAL PRIMARY KEY,
    list_id  INT REFERENCES todo_lists (id) ON DELETE CASCADE,
    item_id  INT REFERENCES todo_items (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE lists_items;
DROP TABLE users_lists;
DROP TABLE todo_lists;
DROP TABLE users;
DROP TABLE todo_items;
