CREATE TABLE transactions
(
    id            SERIAL PRIMARY KEY,
    user_id       INTEGER        NOT NULL,
    user_email    VARCHAR(255)   NOT NULL,
    amount        DECIMAL(27, 4) NOT NULL,
    currency_code VARCHAR(3)     NOT NULL,
    creation_time TIMESTAMP      NOT NULL DEFAULT NOW(),
    modified_time TIMESTAMP      NOT NULL DEFAULT NOW(),
    status        INTEGER        NOT NULL
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