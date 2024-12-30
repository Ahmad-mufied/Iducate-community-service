-- Drop the database if it exists
DROP DATABASE IF EXISTS iducate_dev;

-- Create the database
CREATE DATABASE iducate_dev;

-- Enum Definitions
CREATE TYPE gender AS ENUM ('Man', 'Women');
CREATE TYPE degree AS ENUM ('Diploma', 'Bachelor', 'Master', 'Doctoral');
CREATE TYPE country AS ENUM ('Germany', 'US', 'Malaysia', 'Australia');
CREATE TYPE major AS ENUM ('Art', 'Science', 'Social');

-- Table: Users
DROP TABLE IF EXISTS users;
CREATE TABLE users
(
    id       VARCHAR(100) PRIMARY KEY,
    email    VARCHAR(50) UNIQUE NOT NULL,
    username VARCHAR(100)       NOT NULL,
    gender   gender,
    country  country,
    degree   degree,
    major    major
);

-- Table: Posts
DROP TABLE IF EXISTS posts;
CREATE TABLE posts
(
    id         SERIAL PRIMARY KEY,
    user_id    VARCHAR(100)                NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    title      VARCHAR(255)                NOT NULL,
    content    TEXT                        NOT NULL,
    views      INT                                  DEFAULT 0,
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(), -- Creation timestamp
    updated_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()  -- Last update timestamp
);


-- Table: Comments
DROP TABLE IF EXISTS comments;
CREATE TABLE comments
(
    id         SERIAL PRIMARY KEY,
    post_id    INT                         NOT NULL REFERENCES posts (id) ON DELETE CASCADE,
    user_id    VARCHAR(100)                         NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    content    TEXT                        NOT NULL,
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(), -- Creation timestamp
    updated_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()  -- Last update timestamp
);

-- Table: Likes
DROP TABLE IF EXISTS likes;
CREATE TABLE likes
(
    id         SERIAL PRIMARY KEY,
    user_id    VARCHAR(100)                         NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    post_id    INT                         NOT NULL REFERENCES posts (id) ON DELETE CASCADE,
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- Create an index on id, email, and username columns
CREATE INDEX idx_users_id_email_username ON users (id, email, username);

-- Create individual indexes on id, email, and username columns
CREATE INDEX idx_users_id ON users (id);
CREATE INDEX idx_users_email ON users (email);
CREATE INDEX idx_users_username ON users (username);

-- Drop individual indexes on id, email, and username columns
DROP INDEX IF EXISTS idx_users_id;
DROP INDEX IF EXISTS idx_users_email;
DROP INDEX IF EXISTS idx_users_username;

-- Drop the composite index on id, email, and username columns
DROP INDEX IF EXISTS idx_users_id_email_username;
