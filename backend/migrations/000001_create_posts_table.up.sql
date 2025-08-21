CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE posts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    platform TEXT NOT NULL,
    platform_post_id TEXT NOT NULL,
    message TEXT NOT NULL,
    media_url TEXT,
    posted_at TIMESTAMP WITH TIME ZONE NOT NULL,
    status TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE INDEX idx_posts_user_id ON posts(user_id);
