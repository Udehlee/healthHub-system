 CREATE TABLE users (
    user_id BIGSERIAL PRIMARY KEY,
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    pass_word TEXT NOT NULL,
    user_role TEXT,
    gender TEXT,
    user_address TEXT
);

