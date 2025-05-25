-- +migrate Up
CREATE TABLE mailings (
    id SERIAL PRIMARY KEY,
    message TEXT NOT NULL,
    group_name VARCHAR(255) NOT NULL,
    send_time TIMESTAMP NOT NULL,
    periodicity VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE groups (
    name VARCHAR(255) PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE group_users (
    group_name VARCHAR(255) REFERENCES groups(name),
    telegram_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (group_name, telegram_id)
);

-- +migrate Down
DROP TABLE group_users;
DROP TABLE groups;
DROP TABLE mailings;