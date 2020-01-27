-- Create tables for rconfig software

CREATE TABLE macs (
    mac_id SERIAL PRIMARY KEY,
    mac VARCHAR(17)
);

CREATE TABLE configs (
    config_id SERIAL PRIMARY KEY,
    ssid VARCHAR(32) NOT NULL,
    lease_time VARCHAR(4)
);

CREATE TABLE configs_macs (
    config_id INT REFERENCES configs (config_id),
    mac_id INT REFERENCES macs (mac_id),
    PRIMARY KEY (config_id, mac_id)
);

CREATE TABLE boxes (
    box_id SERIAL PRIMARY KEY,
    boxname VARCHAR(128) NOT NULL,
    config_id INT REFERENCES configs (config_id)
);

CREATE TABLE owners (
    owner_id SERIAL PRIMARY KEY,
    owner VARCHAR(256) NOT NULL,
    email VARCHAR(64),
    box_id INT REFERENCES boxes (box_id)
);


CREATE TABLE owner_boxes (
    owner_id INT REFERENCES owners (owner_id),
    box_id INT REFERENCES boxes (box_id),
    PRIMARY KEY (owner_id, box_id)
);
