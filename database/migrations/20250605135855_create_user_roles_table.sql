-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_roles (
                          user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                          role_id UUID NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
                          created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
                          PRIMARY KEY (user_id, role_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_roles;
-- +goose StatementEnd
