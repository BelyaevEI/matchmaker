-- +goose Up
-- +goose StatementBegin
CREATE TABLE User
(
    name VARCHAR(255) PRIMARY KEY NOT NULL,
    skill int
    latency int
    created_at TIMESTAMP NOT NULL DEFAULT now(),
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE User
-- +goose StatementEnd
