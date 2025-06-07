-- +goose Up
-- +goose StatementBegin
CREATE TABLE permissions (
                           id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                           group_id UUID REFERENCES permission_groups(id) ON DELETE SET NULL,
                           name VARCHAR(100) NOT NULL UNIQUE,
                           label VARCHAR(255) NOT NULL,
                           description TEXT,
                           created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
                           updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS permissions;
-- +goose StatementEnd
