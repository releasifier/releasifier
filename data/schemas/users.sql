DROP TABLE IF EXISTS users;
DROP SEQUENCE IF EXISTS user_id_seq;

CREATE SEQUENCE user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE users (
    id bigint DEFAULT nextval('user_id_seq'::regclass) NOT NULL,
    fullname varchar(128) NOT NULL,
    email varchar(256) NOT NULL,
    password varchar(128) NOT NULL,
    created_at timestamp DEFAULT now() NOT NULL
);

ALTER TABLE ONLY users ADD CONSTRAINT users_pkey PRIMARY KEY (id);
ALTER TABLE users ADD UNIQUE ("email");

-- # INSERT Root user
-- # please change this according to your specification
INSERT INTO users (id, fullname, email, password) VALUES (1, 'Mr. Robot', 'root', 'root');
