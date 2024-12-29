-- +goose Up
-- +goose StatementBegin
INSERT INTO rooms (number, type, price)
VALUES ('101', 1, 5000),
       ('102', 1, 5000),
       ('201', 2, 8000),
       ('202', 2, 8000),
       ('301', 3, 15000);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM rooms WHERE number in ('101', '102', '201', '202', '301');
-- +goose StatementEnd
