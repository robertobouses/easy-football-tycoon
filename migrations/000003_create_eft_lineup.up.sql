BEGIN;

CREATE TABLE eft.lineup (
    playerid UUID PRIMARY KEY,
    FOREIGN KEY (playerid) REFERENCES eft.team(playerid)
);

COMMIT;