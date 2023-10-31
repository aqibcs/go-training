BEGIN;
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100)
);

CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    content TEXT
);

INSERT INTO users (name, email)
VALUES ('Aqib', 'aqib@gmail.com'),
       ('Asad', 'asad@gmail.com'),
       ('Nauman', 'nauman@gmail.com');

INSERT INTO posts (user_id, content)
VALUES (1, 'Post by Aqib'),
       (2, 'Post by Asad'),
       (3, 'Post by Nauman');

COMMIT;
