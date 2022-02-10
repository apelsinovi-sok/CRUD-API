-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id        text NOT NULL,
    firstname text NOT NULL,
    lastname  text NOT NULL,
    email     text NOT NULL,
    age       int  NOT NULL,
    created   TIMESTAMP NOT NULL,
    PRIMARY KEY (id)
)

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
