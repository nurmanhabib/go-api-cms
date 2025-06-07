-- +goose Up
-- +goose StatementBegin
CREATE TABLE permission_groups (
                                 id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                                 name VARCHAR(100) NOT NULL UNIQUE,
                                 label VARCHAR(255) NOT NULL,
                                 description TEXT,
                                 created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
                                 updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS permission_groups;
-- +goose StatementEnd
