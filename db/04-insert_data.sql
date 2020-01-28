-- Insert data into "configs" table
INSERT INTO configs (ssid, lease_time)
VALUES ('ssid1', '30m'),
       ('ssid2', '30m'),
       ('ssid3', NULL),
       ('ssid4', '1h');

-- Insert data into "macs" table
INSERT INTO macs (mac, config_id)
VALUES ('11:11:11:11:11:11', 1),
       ('22:22:22:22:22:22', 1),
       ('33:33:33:33:33:33', 3);

-- Insert data into "owners" table
INSERT INTO owners (owner, email)
VALUES ('Owner 1', 'owner1@owner1.com'),
       ('Owner 2', 'owner2@owner2.com');

-- Insert data into "boxes" table
INSERT INTO boxes (boxname, config_id, owner_id)
VALUES ('box1', 1, 1),
       ('box2', 1, 1),
       ('box3', 3, 1),
       ('box4', 4, 2);
