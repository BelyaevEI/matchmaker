-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    name VARCHAR(255) PRIMARY KEY NOT NULL,
    skill INT,
    latency INT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
