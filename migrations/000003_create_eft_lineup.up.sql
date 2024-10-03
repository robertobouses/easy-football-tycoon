BEGIN;

CREATE TABLE eft.lineup (
    playerid UUID PRIMARY KEY NOT NULL, 
    lastname VARCHAR(255) NOT NULL,
    position VARCHAR(255) NOT NULL,
    technique INT,
    mental INT,
    physique INT
);

COMMIT;