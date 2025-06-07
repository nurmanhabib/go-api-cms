-- +goose Up
-- +goose StatementBegin
CREATE TABLE roles (
                     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                     name VARCHAR(100) NOT NULL UNIQUE,
                     description TEXT,
                     is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
                     created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
                     updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS roles;
-- +goose StatementEnd
