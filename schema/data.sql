BEGIN;

INSERT INTO users (name, email)
VALUES ('Aqib', 'aqib@gmail.com'),
       ('Asad', 'asad@gmail.com'),
       ('Nauman', 'nauman@gmail.com');

INSERT INTO posts (user_id, content)
VALUES (1, 'Post by Aqib'),
       (2, 'Post by Asad'),
       (3, 'Post by Nauman');

COMMIT;
