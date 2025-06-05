-- +goose Up
-- +goose StatementBegin
CREATE TABLE role_permissions (
                                role_id UUID NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
                                permission_id UUID NOT NULL REFERENCES permissions(id) ON DELETE CASCADE,
                                created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
                                PRIMARY KEY (role_id, permission_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS role_permissions;
-- +goose StatementEnd
