-- +goose Up
CREATE TABLE tbl_users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(20) DEFAULT NULL,
    last_name VARCHAR(20) DEFAULT NULL,
    user_name VARCHAR(40) DEFAULT NULL UNIQUE,
    password VARCHAR(100) DEFAULT NULL,
    email VARCHAR(25) DEFAULT NULL,
    phone VARCHAR(15) DEFAULT NULL,
    is_admin SMALLINT DEFAULT 0,
    login_session VARCHAR(350) DEFAULT NULL,
    role_id INTEGER DEFAULT 0,
    currency_id INTEGER DEFAULT 1,
    lang VARCHAR(20) DEFAULT 'km',
    session_share VARCHAR(350) DEFAULT NULL,
    status_id SMALLINT DEFAULT 1,
    "order" INTEGER DEFAULT 0,
    created_by INTEGER,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_by INTEGER,
    updated_at TIMESTAMP WITHOUT TIME ZONE,
    deleted_by INTEGER,
    deleted_at TIMESTAMP WITHOUT TIME ZONE

);

-- +goose StatementBegin
INSERT INTO tbl_users (
    first_name, last_name, user_name, password, email, phone, is_admin, role_id, currency_id, lang, status_id, created_by
) VALUES 
('ADMIN', 'SUPER1', 'ADMIN1', '123456', 'ADMIN1@example.com', '1234567890', 1, 1, 1, 'km', 1, 1),
('ADMIN', 'SUPER2', 'ADMIN2', '123456', 'ADMIN1@example.com', '1234567890', 1, 1, 1, 'km', 1, 1),
('ADMIN', 'SUPER14', 'ADMIN14', '123456', 'ADMIN2@example.com', '0987654321', 0, 1, 1, 'km', 1, 1);
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS tbl_users;
