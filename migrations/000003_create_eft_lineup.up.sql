BEGIN;

CREATE TABLE eft.lineup (
    playerid UUID PRIMARY KEY NOT NULL, 
    playername VARCHAR(255) NOT NULL,
    position VARCHAR(255) NOT NULL,
    technique INT,
    mental INT,
    physique INT
);

COMMIT;


-- BEGIN;

-- CREATE TABLE eft.lineup (
--     playerid UUID NOT NULL PRIMARY KEY,
--     playername VARCHAR(255) NOT NULL,
--     position VARCHAR(255) NOT NULL,
--     technique INT,
--     mental INT,
--     physique INT,
--     FOREIGN KEY (playerid) REFERENCES eft.team(playerid)
-- );



-- COMMIT;