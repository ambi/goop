DROP DATABASE IF EXISTS goop;
CREATE DATABASE goop;
USE goop;

CREATE TABLE clients (
    uuid CHAR(36) NOT NULL PRIMARY KEY,
    name VARCHAR(1000) NOT NULL,
    client_id VARCHAR(100) NOT NULL, -- TODO: add index
    client_secret VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE users (
    uuid CHAR(36) NOT NULL PRIMARY KEY,
    login_id VARCHAR(1000) NOT NULL, -- TODO: add index
    email VARCHAR(1000) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE redirect_uris (
    uuid CHAR(36) NOT NULL PRIMARY KEY,
    client_uuid CHAR(36) NOT NULL,
    uri VARCHAR(1000) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_client_uuid
      FOREIGN KEY (client_uuid)
      REFERENCES clients(uuid)
      ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE authorization_codes (
    uuid CHAR(36) NOT NULL PRIMARY KEY,
    user_uuid CHAR(36) NOT NULL,
    code VARCHAR(1000) NOT NULL, -- TODO: add index
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_user_uuid
      FOREIGN KEY (user_uuid)
      REFERENCES users(uuid)
      ON DELETE CASCADE ON UPDATE CASCADE
);

SET @client1 = UUID();
SET @user1 = UUID();
INSERT INTO clients (uuid, name, client_id, client_secret) VALUES (@client1, "client1", "client-id1", "client-secret1");
INSERT INTO redirect_uris (uuid, client_uuid, uri) VALUES (UUID(), @client1, "http://localhost/callback");
INSERT INTO users (uuid, login_id, email) VALUES (@user1, "test1@example.com", "test1@example.jp");
