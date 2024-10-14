BEGIN;

CREATE TABLE eft.analytics (
    analyticsid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    training INT,
    finances INT,
    scouting INT,
    physiotherapy INT,
    totalsalaries INT,
    totaltraining INT,
    totalfinances INT,
    totalscouting INT,
    totalphysiotherapy INT,
    trainingstaffcount INT,
    financestaffcount INT,
    scoutingstaffcount INT,
    physiotherapystaffcount INT,
    trust INT,
    stadiumcapacity INT,
    job VARCHAR(255)
);

COMMIT;
