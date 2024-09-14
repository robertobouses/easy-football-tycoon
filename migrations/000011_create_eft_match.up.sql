BEGIN;

CREATE TABLE eft.match (
matchid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
calendaryid INT,
matchdaynumber INT,
homeoraway  VARCHAR(255) NOT NULL,
rivalname  VARCHAR(255) NOT NULL,
teamballpossession INT,
teamscoringchances INT,
teamgoals INT,
rivalballpossession INT,
rivalscoringchances INT,
rivalgoals INT
);

COMMIT;