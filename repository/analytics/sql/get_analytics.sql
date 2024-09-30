SELECT
    t.analyticsid,
    t.training,
    t.finances,
    t.scouting,
    t.physiotherapy,
    t.totalsalaries,
    t.totaltraining,
    t.totalfinances,
    t.totalscouting,
    t.totalphysiotherapy,
    COUNT(t.training) OVER () AS trainingstaffcount,
    COUNT(t.finances) OVER () AS financestaffcount,
    COUNT(t.scouting) OVER () AS scoutingstaffcount,
    COUNT(t.physiotherapy) OVER () AS physiotherapystaffcount,
    t.trust,
    t.stadiumcapacity
FROM
    eft.analytics t
ORDER BY
    t.analyticsid DESC  
LIMIT 1;