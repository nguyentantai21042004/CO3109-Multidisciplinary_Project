CREATE TABLE IF NOT EXISTS uploads (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    path TEXT NOT NULL,
    source VARCHAR(255) NOT NULL,
    from_location VARCHAR(255) NOT NULL,
    public_id VARCHAR(255),
    created_user_id UUID NOT NULL,  
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_uploads_created_user_id ON uploads(created_user_id);
CREATE INDEX idx_uploads_public_id ON uploads(public_id);

ALTER TABLE uploads ADD FOREIGN KEY (created_user_id) REFERENCES users(id);