BEGIN;

CREATE TABLE eft.prospect (
prospectid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
prospectname VARCHAR(255) NOT NULL,
position VARCHAR(255) NOT NULL,
age INT,
fee INT,
salary INT,
technique INT,
mental INT,
physique INT,
injurydays INT DEFAULT 0,
job VARCHAR(255) NOT NULL,
rarity INT
);

COMMIT;