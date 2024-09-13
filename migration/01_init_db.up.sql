CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "users" (
    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "username" VARCHAR(255) UNIQUE,
    "first_name" VARCHAR(255),
    "last_name" VARCHAR(255),
    "email" VARCHAR(255) UNIQUE NOT NULL,
    "password_hash" TEXT NOT NULL, 
    "created_at" INT NOT NULL
);

CREATE TABLE tweets (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    media_url TEXT, -- optional media (image/video)
    created_at INT NOT NULL
);

CREATE TABLE followers (
    id SERIAL PRIMARY KEY,
    follower_id INT REFERENCES users(id) ON DELETE CASCADE,  -- Kuzatuvchi foydalanuvchi IDsi
    following_id INT REFERENCES users(id) ON DELETE CASCADE, -- Kuzatilayotgan foydalanuvchi IDsi
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP            -- Kuzatish qachon amalga oshirilganligi
);


CREATE TABLE likes (
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    tweet_id INT REFERENCES tweets(id) ON DELETE CASCADE,
    created_at INT NOT NULL,
    PRIMARY KEY(user_id, tweet_id)
);

CREATE TABLE retweets (
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    tweet_id INT REFERENCES tweets(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(user_id, tweet_id)
);

