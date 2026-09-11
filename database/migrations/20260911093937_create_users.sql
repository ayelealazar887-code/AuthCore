-- +goose Up

CREATE TABLE users (
    id CHAR(36) NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(30),
    token VARCHAR(255),
    user_type ENUM('admin', 'employee') NOT NULL DEFAULT 'employee',
    refresh_token VARCHAR(255),
    user_id CHAR(36),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (id)
);

-- +goose Down

DROP TABLE users;