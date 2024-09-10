SELECT
    t.analyticsid,
    t.training,
    t.finances,
    t.scouting,
    t.physiotherapy,
    t.totalfinances,
    t.totalscouting,
    t.totalphysiotherapy,
    t.totaltraining,
    t.totalsalaries,
    COUNT(t.training) OVER () AS trainingstaffcount,
    COUNT(t.finances) OVER () AS financestaffcount,
    COUNT(t.scouting) OVER () AS scoutingstaffcount,
    COUNT(t.physiotherapy) OVER () AS physiotherapystaffcount
FROM
    eft.analytics t;
