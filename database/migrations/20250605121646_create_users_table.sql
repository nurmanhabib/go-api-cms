-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
                     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                     name VARCHAR(255) NOT NULL,
                     username VARCHAR(255) NOT NULL UNIQUE,
                     email VARCHAR(255) NOT NULL UNIQUE,
                     password TEXT NOT NULL,
                     is_active BOOLEAN NOT NULL DEFAULT TRUE,
                     is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
                     created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
                     updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
