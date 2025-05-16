CREATE TABLE eft.player_stats (
    playerid UUID PRIMARY KEY,
    appearances INT DEFAULT 0,
    blocks INT DEFAULT 0,
    saves INT DEFAULT 0,
    aerialduel INT DEFAULT 0,
    keypass INT DEFAULT 0,
    assists INT DEFAULT 0,
    chances INT DEFAULT 0,
    goals INT DEFAULT 0,
    mvp INT DEFAULT 0,
    rating FLOAT DEFAULT 0.0,
    FOREIGN KEY (playerid) REFERENCES eft.team (playerid)
);
