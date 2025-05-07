CREATE TABLE users (
    user_id       BIGSERIAL PRIMARY KEY,
    firstname     VARCHAR(100) NOT NULL,
    lastname      VARCHAR(100) NOT NULL,
    email         VARCHAR(255) UNIQUE NOT NULL,
    pass_word     TEXT NOT NULL,
    gender        VARCHAR(10),
    user_address  TEXT,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
