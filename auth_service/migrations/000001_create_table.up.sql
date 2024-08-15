CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(230) NOT NULL,
    email VARCHAR(230) NOT NULL unique,
    password VARCHAR(230) NOT NULL,
    full_name VARCHAR(230) not NULL,
    dob DATE,
    role VARCHAR(90),
    language VARCHAR,
    confirmed bool,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at INT DEFAULT 0
);
create TABLE IF NOT EXISTS notifications (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    reciever_id UUID NOT NULL,
    sender_id UUID,
    message TEXT NOT NULL,
    status  VARCHAR(50) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at INT DEFAULT 0
)