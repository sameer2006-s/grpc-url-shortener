CREATE TABLE links (
    id UUID PRIMARY KEY,
    short_code TEXT NOT NULL UNIQUE,
    url TEXT NOT NULL,
    clicks INTEGER DEFAULT 0,
    created_at TIMESTAMP NOT NULL
);