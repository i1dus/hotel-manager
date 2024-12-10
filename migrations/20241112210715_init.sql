-- +goose Up
-- +goose StatementBegin
CREATE TABLE rooms
(
    id      SERIAL PRIMARY KEY,
    number  TEXT    NOT NULL UNIQUE,
    type    INTEGER NOT NULL,
    price   INTEGER NOT NULL,
    cleaned BOOLEAN NOT NULL DEFAULT true
);

CREATE TABLE employees
(
    id       SERIAL PRIMARY KEY,
    username TEXT    NOT NULL,
    name     TEXT,
    position INTEGER NOT NULL
);

CREATE TABLE clients
(
    id       SERIAL PRIMARY KEY,
    name     TEXT NOT NULL,
    surname  TEXT NOT NULL,
    passport TEXT NOT NULL UNIQUE
);

ALTER TABLE employees
    ADD CONSTRAINT unique_user_position UNIQUE (username, position);

CREATE TABLE room_occupancies
(
    id          SERIAL PRIMARY KEY,
    room_number TEXT        NOT NULL,
    passport    TEXT        NOT NULL,
    start_at    timestamptz NOT NULL,
    end_at      timestamptz,
    description TEXT        NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE rooms;
DROP TABLE employees;
DROP TABLE clients;
DROP TABLE room_occupancies;
-- +goose StatementEnd
