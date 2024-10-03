SELECT
    lastname,
    position,
    technique,
    mental,
    physique,
    SUM(l.technique) OVER () AS total_technique,
    SUM(l.mental) OVER () AS total_mental,
    SUM(l.physique) OVER () AS total_physique
FROM
    eft.lineup l;