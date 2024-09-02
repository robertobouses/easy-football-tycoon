BEGIN;

CREATE TABLE eft.analytics (
analyticsid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
finances INT,
scouting INT,
physiotherapy INT
);

COMMIT;