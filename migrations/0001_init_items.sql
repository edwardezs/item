-- +goose Up
CREATE TABLE items
(
    id              SERIAL PRIMARY KEY,
    title           VARCHAR(255) NOT NULL,
    description     VARCHAR(255),
    done            BOOLEAN NOT NULL DEFAULT FALSE
);

INSERT INTO items (title, description, done) VALUES 
    ('Buy milk', 'Buy 2 liters of milk', FALSE), 
    ('Walk the dog', 'Take the dog for a 30-minute walk', FALSE), 
    ('Do laundry', 'Wash, dry, and fold the clothes', FALSE), 
    ('Clean the room', 'Vacuum and dust the living room', FALSE), 
    ('Call mom', 'Call mom to catch up', FALSE);

-- +goose Down
DROP TABLE items;
