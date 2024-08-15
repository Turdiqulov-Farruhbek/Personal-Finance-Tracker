CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(230) NOT NULL,
    email VARCHAR(230) NOT NULL unique,
    password VARCHAR(500) NOT NULL,
    full_name VARCHAR(230) NOT NULL,
    dob DATE,
    role VARCHAR(90),
    language VARCHAR,
    confirmed bool,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at INT DEFAULT 0
);