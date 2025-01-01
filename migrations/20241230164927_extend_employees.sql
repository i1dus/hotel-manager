-- +goose Up
-- +goose StatementBegin
ALTER TABLE employees
    ADD COLUMN user_id INTEGER;

ALTER TABLE employees
    ADD UNIQUE (username);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE employees
    DROP COLUMN user_id;
-- +goose StatementEnd
