SELECT
    t.playername,
    t.position,
    t.technique,
    t.mental,
    t.physique,
    SUM(t.technique) OVER () AS total_technique,
    SUM(t.mental) OVER () AS total_mental,
    SUM(t.physique) OVER () AS total_physique
FROM
    eft.lineup l
JOIN
    eft.team t ON l.playerid = t.playerid;
