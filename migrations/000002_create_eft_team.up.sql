BEGIN;

CREATE TABLE eft.team (
    playerid UUID PRIMARY KEY,
    playername VARCHAR(255) NOT NULL,
    position VARCHAR(255) NOT NULL,
    age INT,
    fee INT,
    salary INT,
    technique INT,
    mental INT,
    physique INT,
    injurydays INT,
    lined BOOLEAN,
    PRIMARY KEY (playerid)
);

COMMIT;