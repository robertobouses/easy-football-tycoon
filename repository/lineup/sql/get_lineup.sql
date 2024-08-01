SELECT
    t.playername,
    t.position,
    t.technique,
    t.mental,
    t.physique
FROM
    eft.lineup l
JOIN
    eft.team t ON l.playerid = t.playerid;
