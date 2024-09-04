SELECT
    t.analyticsid,
    t.finances,
    t.scouting,
    t.physiotherapy,
    SUM(t.finances) OVER () AS total_finances,
    SUM(t.scouting) OVER () AS total_scouting,
    SUM(t.physiotherapy) OVER () AS total_physiotherapy
FROM
    eft.analytics t;