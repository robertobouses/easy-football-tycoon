BEGIN;

CREATE TABLE eft.player_stats (
    player_id UUID PRIMARY KEY,
    appearances INT DEFAULT 0,
    chances INT DEFAULT 0,
    assists INT DEFAULT 0,
    goals INT DEFAULT 0,
    mvp INT DEFAULT 0,
    rating FLOAT DEFAULT 0.0,
    FOREIGN KEY (player_id) REFERENCES eft.team (playerid)
)

COMMIT;
