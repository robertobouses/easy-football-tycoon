BEGIN;

CREATE TABLE eft.rivals (
    teamid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    teamname VARCHAR(255) NOT NULL,
    technique INT,
    mental INT,
    physique INT,
    
);
