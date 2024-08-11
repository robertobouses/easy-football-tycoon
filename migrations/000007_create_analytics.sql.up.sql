BEGIN;

CREATE TABLE eft.analytics (
analyticsid UUID PRIMARY KEY DEFAULT,
finances INT,
scouting INT,
physiotherapy INT
);

COMMIT;