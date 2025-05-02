CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    created_at TIMESTAMP default NOW(),
    updated_at TIMESTAMP default NULL,
    deleted_at TIMESTAMP default NULL
)