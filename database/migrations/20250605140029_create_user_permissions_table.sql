-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_permissions (
                                user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                                permission_id UUID NOT NULL REFERENCES permissions(id) ON DELETE CASCADE,
                                created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
                                PRIMARY KEY (user_id, permission_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_permissions;
-- +goose StatementEnd
