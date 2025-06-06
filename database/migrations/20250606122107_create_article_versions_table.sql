-- +goose Up
-- +goose StatementBegin
CREATE TABLE article_versions (
                                id UUID PRIMARY KEY,
                                article_id UUID NOT NULL REFERENCES articles(id) ON DELETE CASCADE,
                                version INTEGER NOT NULL,
                                title TEXT NOT NULL,
                                content TEXT NOT NULL,
                                status VARCHAR(20) NOT NULL CHECK (status IN ('draft', 'published', 'archived')),
                                created_by UUID,
                                is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
                                created_at TIMESTAMPTZ DEFAULT NOW(),

                                CONSTRAINT unique_article_version UNIQUE (article_id, version)
);

CREATE INDEX idx_article_versions_article_id ON article_versions(article_id);
CREATE INDEX idx_article_versions_status ON article_versions(status);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS article_versions;
-- +goose StatementEnd
