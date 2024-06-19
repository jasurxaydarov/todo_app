CREATE DATABASE todo;

CREATE TABLE users(
    user_id         UUID               PRIMARY KEY  DEFAULT gen_random_uuid(),
    user_name       VARCHAR(124)       UNIQUE       NOT NULL,
    password        VARCHAR(69)                     NOT NULL,
    created_at      TIMESTAMP                       DEFAULT NOW()    
);


CREATE TABLE todo(
    todo_id         UUID               PRIMARY KEY  DEFAULT gen_random_uuid(),
    title           VARCHAR(255)                    NOT NULL,
    descreption     TEXT,
    is_completed    BOOLEAN                         DEFAULT FALSE,
    created_at      TIMESTAMP                       DEFAULT NOW(),                   
    user_id         UUID                            REFERENCES  users(user_id)
);
