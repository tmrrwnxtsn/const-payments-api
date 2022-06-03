CREATE TABLE users
(
    id    INTEGER PRIMARY KEY,
    email VARCHAR(255) NOT NULL
);

CREATE TABLE transactions
(
    id            SERIAL PRIMARY KEY,
    user_id       INTEGER        NOT NULL,
    amount        DECIMAL(27, 4) NOT NULL,
    currency_code VARCHAR(3)     NOT NULL,
    creation_time TIMESTAMP      NOT NULL DEFAULT NOW(),
    modified_time TIMESTAMP      NOT NULL DEFAULT NOW(),
    status        INTEGER        NOT NULL DEFAULT 0,
    CONSTRAINT fk_user_transaction FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE FUNCTION trigger_set_timestamp()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.modified_time = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp
    BEFORE UPDATE
    ON transactions
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();