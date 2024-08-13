BEGIN;

CREATE TABLE eft.analytics (
analyticsid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
finances INT,
scouting INT,
physiotherapy INT
);

COMMIT;