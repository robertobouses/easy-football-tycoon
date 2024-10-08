INSERT INTO eft.analytics (
    training,
    finances,
    scouting,
    physiotherapy,
    totalsalaries,
    totaltraining,
    totalfinances,
    totalscouting,
    totalphysiotherapy,
    trainingstaffcount, 
    financestaffcount,
    scoutingstaffcount,
    physiotherapystaffcount,
    trust,
    stadiumcapacity,
    job
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16);
