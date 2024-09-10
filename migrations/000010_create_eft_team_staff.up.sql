BEGIN;

CREATE TABLE eft.team_staff (
staffid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
staffname VARCHAR(255) NOT NULL,
job VARCHAR(255) NOT NULL,
age INT,
fee INT,
salary INT,
training INT,
finances INT,
scouting INT,
physiotherapy INT,
rarity INT
);

COMMIT;