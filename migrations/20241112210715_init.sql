-- +goose Up
-- +goose StatementBegin
CREATE TABLE rooms (
    id SERIAL PRIMARY KEY,
    in_use BOOLEAN
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE rooms;
-- +goose StatementEnd
