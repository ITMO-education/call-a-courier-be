-- Thank you for giving goose a try!

-- +goose Up
CREATE TABLE IF NOT EXISTS contracts
(
    address VARCHAR(48) PRIMARY KEY,
    owner   VARCHAR(48)
);

-- +goose Down
DROP TABLE IF EXISTS contracts;
