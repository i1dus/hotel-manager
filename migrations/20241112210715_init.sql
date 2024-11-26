-- +goose Up
-- +goose StatementBegin
CREATE TABLE rooms
(
    id     SERIAL PRIMARY KEY,
    in_use BOOLEAN NOT NULL DEFAULT false,
    type   INTEGER NOT NULL
);

CREATE TABLE employees
(
    id       SERIAL PRIMARY KEY,
    username TEXT    NOT NULL,
    name     TEXT,
    position INTEGER NOT NULL
);

ALTER TABLE employees ADD CONSTRAINT unique_user_position UNIQUE (username, position);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE rooms;
DROP TABLE employees;
-- +goose StatementEnd
