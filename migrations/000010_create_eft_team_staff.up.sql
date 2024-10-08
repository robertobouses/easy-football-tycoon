BEGIN;

CREATE TABLE eft.team_staff (
staffid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
firstname VARCHAR(255) NOT NULL,
lastname VARCHAR(255) NOT NULL,
nationality VARCHAR(255) NOT NULL,
job VARCHAR(255) NOT NULL,
age INT,
fee INT,
salary INT,
training INT,
finances INT,
scouting INT,
physiotherapy INT,
knowledge INT,
intelligence INT,
rarity INT
);

COMMIT;