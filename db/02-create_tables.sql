-- Create tables for rconfig software

CREATE TABLE configs (
    config_id SERIAL PRIMARY KEY,
    ssid VARCHAR(32) NOT NULL,
    lease_time VARCHAR(4)
);

CREATE TABLE macs (
    mac_id SERIAL PRIMARY KEY,
    mac VARCHAR(17) UNIQUE,
    config_id INT REFERENCES configs (config_id)
);

CREATE TABLE owners (
    owner_id SERIAL PRIMARY KEY,
    owner VARCHAR(256) NOT NULL,
    email VARCHAR(64)
);

CREATE TABLE boxes (
    box_id SERIAL PRIMARY KEY,
    boxname VARCHAR(128) NOT NULL,
    config_id INT REFERENCES configs (config_id),
    owner_id INT REFERENCES owners (owner_id)
);
