INSERT INTO eft.match (
    calendaryid,
    matchdaynumber,
    homeoraway,
    rivalname,
    teamballpossession,
    teamscoringchances,
    teamgoals,
    rivalballpossession,
    rivalscoringchances,
    rivalgoals
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
);
