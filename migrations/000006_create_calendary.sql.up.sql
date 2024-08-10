BEGIN;

CREATE TABLE eft.calendary (
    calendaryid SERIAL PRIMARY KEY,
    daytype VARCHAR(255) NOT NULL
);


COMMIT;
