CREATE TABLE  users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE,
    full_name VARCHAR(255) ,
    bio TEXT NOT NULL,
    profile_image_url TEXT ,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL, 
    created_at INT NOT NULL
);

CREATE TABLE tweets (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    media_url TEXT, -- optional media 
    likes_count INT DEFAULT 0,
    retweets_count INT DEFAULT 0 ,
    created_at INT NOT NULL
);

CREATE TABLE followers (
    id SERIAL PRIMARY KEY,
    follower_id INT REFERENCES users(id) ON DELETE CASCADE,  
    following_id INT REFERENCES users(id) ON DELETE CASCADE,
    created_at INT NOT NULL
);


CREATE TABLE likes (
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    tweet_id INT REFERENCES tweets(id) ON DELETE CASCADE,
    created_at INT NOT NULL,
    PRIMARY KEY(user_id, tweet_id)
);

CREATE TABLE retweets (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    tweet_id INT REFERENCES tweets(id) ON DELETE CASCADE,
    created_at INT NOT NULL,
    UNIQUE(user_id, tweet_id)
);