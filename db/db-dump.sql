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

-- Seed data
-- don@email.com password
-- adel@email.com hahahihi
-- ship@email.com ship123
-- myadmin@email.com admin123
-- brad@email.com brad456

CREATE TABLE transactions(
    id SERIAL PRIMARY KEY,
    amount BIGINT,
    source_wallet_id INT,
    destination_wallet_id INT,
    description VARCHAR(35),
    transaction_type VARCHAR(12),
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp without time zone,
    source_of_fund_id INT,
    CONSTRAINT fk_source_wallet FOREIGN KEY (source_wallet_id) REFERENCES wallets(id),
    CONSTRAINT fk_destination_wallet FOREIGN KEY (destination_wallet_id) REFERENCES wallets(id),
    CONSTRAINT fk_source_of_fund FOREIGN KEY (source_of_fund_id) REFERENCES source_of_funds(id)
);