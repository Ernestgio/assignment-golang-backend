CREATE TABLE source_of_funds(
	id SERIAL PRIMARY KEY,
	name VARCHAR NOT NULL
);

INSERT INTO source_of_funds (name) VALUES ('Bank Transfer');
INSERT INTO source_of_funds (name) VALUES ('Credit Card');
INSERT INTO source_of_funds (name) VALUES ('Cash');



CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    email VARCHAR UNIQUE NOT NULL,
    password VARCHAR,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp without time zone
);

CREATE SEQUENCE wallet_id_seq START 777000 INCREMENT 1;

CREATE TABLE wallets(
    id SERIAL PRIMARY KEY,
    amount BIGINT,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp without time zone,
    user_id INT,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id)
);

ALTER TABLE wallets
ALTER COLUMN id SET DEFAULT nextval('wallet_id_seq');

ALTER SEQUENCE wallet_id_seq
OWNED BY wallets.id;
