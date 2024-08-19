BEGIN;

CREATE TABLE eft.team (
    playerid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    playername VARCHAR(255) NOT NULL,
    position VARCHAR(255) NOT NULL,
    age INT,
    fee INT,
    salary INT,
    technique INT,
    mental INT,
    physique INT,
    injurydays INT DEFAULT 0,
    lined BOOLEAN DEFAULT false

);

COMMIT;
