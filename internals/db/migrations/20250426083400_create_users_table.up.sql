CREATE TABLE users (
    user_id     BIGSERIAL PRIMARY KEY,
    first_name  VARCHAR(100) NOT NULL,
    last_name   VARCHAR(100) NOT NULL,
    email       VARCHAR(255) UNIQUE NOT NULL,
    pass_word    TEXT NOT NULL,
    role_id INT NOT NULL,
    gender      VARCHAR(10),
    user_address     TEXT,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
     FOREIGN KEY (role_id) REFERENCES roles(role_id)
);
