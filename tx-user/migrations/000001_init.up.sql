CREATE SCHEMA tv1;

CREATE TABLE IF NOT EXISTS tv1.users (
    id INT GENERATED ALWAYS AS IDENTITY             NOT NULL,
    email text            NOT NULL CHECK (length(email) < 100)  UNIQUE,
    hash_password text    NOT NULL CHECK (length(hash_password) < 300)
);




