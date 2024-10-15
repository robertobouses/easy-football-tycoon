BEGIN;

CREATE TABLE eft.rivals (
    teamid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    rivalname VARCHAR(255) NOT NULL,
    technique INT,
    mental INT,
    physique INT
);


COMMIT;
