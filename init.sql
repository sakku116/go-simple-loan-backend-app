-- Schema for Loan Management Database

-- User table
CREATE TABLE User (
    id SERIAL PRIMARY KEY,
    uuid VARCHAR(36) UNIQUE NOT NULL,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    role VARCHAR(255) NOT NULL,
    fullname VARCHAR(255) NOT NULL,
    legalname VARCHAR(255) NOT NULL,
    nik VARCHAR(255) NOT NULL,
    birthplace VARCHAR(255) NOT NULL,
    birthdate VARCHAR(255) NOT NULL, -- Format: DD-MM-YYYY
    current_salary FLOAT NOT NULL,
    current_limit FLOAT NOT NULL,
    ktp_photo VARCHAR(255),
    face_photo VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- RefreshToken table
CREATE TABLE RefreshToken (
    id SERIAL PRIMARY KEY,
    uuid VARCHAR(36) UNIQUE NOT NULL,
    user_id INT NOT NULL REFERENCES User(id) ON DELETE CASCADE,
    user_uuid VARCHAR(36) NOT NULL,
    token VARCHAR(36) UNIQUE NOT NULL,
    used_at TIMESTAMP,
    expired_at TIMESTAMP,
    invalid BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Loan table
CREATE TABLE Loan (
    id SERIAL PRIMARY KEY,
    uuid VARCHAR(36) UNIQUE NOT NULL,
    user_id INT NOT NULL REFERENCES User(id) ON DELETE CASCADE,
    user_uuid VARCHAR(36) NOT NULL,
    asset_name VARCHAR(255) NOT NULL,
    ref_number BIGINT NOT NULL,
    otr FLOAT NOT NULL,
    interest_rate_percentage FLOAT DEFAULT 10 NOT NULL,
    interest_rate FLOAT NOT NULL,
    admin_fee_percentage FLOAT DEFAULT 2 NOT NULL,
    admin_fee FLOAT NOT NULL,
    installment_amount FLOAT NOT NULL,
    total_amount FLOAT NOT NULL,
    term_months INT NOT NULL,
    status VARCHAR(255) DEFAULT 'PENDING',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Seed data
-- Initial Admin User
INSERT INTO User (
    uuid, username, password, email, role, fullname, legalname, nik, birthplace, birthdate, current_salary, current_limit
) VALUES (
    gen_random_uuid()::VARCHAR,
    'admin_username',
    'hashed_admin_password',
    'admin_username@gmail.com',
    'admin',
    'admin_username',
    'admin_username',
    '1234567890123456',
    'Jakarta',
    '11-12-2001',
    10000000,
    1000000
);

-- Initial Normal User
INSERT INTO User (
    uuid, username, password, email, role, fullname, legalname, nik, birthplace, birthdate, current_salary, current_limit
) VALUES (
    gen_random_uuid()::VARCHAR,
    'user_username',
    'hashed_user_password',
    'user_username@gmail.com',
    'user',
    'user_username',
    'user_username',
    '0987654321654321',
    'PATI',
    '11-12-2001',
    11000000,
    100000
);
