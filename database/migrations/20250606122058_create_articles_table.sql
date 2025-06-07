-- +goose Up
-- +goose StatementBegin
CREATE TABLE articles (
                        id UUID PRIMARY KEY,
                        slug VARCHAR(255) UNIQUE NOT NULL,
                        current_version_id UUID,
                        is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
                        created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
                        updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS articles;
-- +goose StatementEnd
