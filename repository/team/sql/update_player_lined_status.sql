UPDATE eft.team
SET lined = CASE WHEN lined THEN FALSE ELSE TRUE END
WHERE playerid = $1;
