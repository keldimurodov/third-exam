CREATE TABLE Users (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    username VARCHAR(32),
    bio VARCHAR(300),
    website VARCHAR(64),
    email VARCHAR(320),
    password VARCHAR(32),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updeted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE Posts (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    userID UUID,
    content TEXT,
    title VARCHAR(256),
    likes INTEGER,
    dislikes INTEGER,
    views INTEGER,
    categories VARCHAR(128),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updeted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE Comments (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    owner_id UUID,
    post_id UUID,
    text TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updeted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);