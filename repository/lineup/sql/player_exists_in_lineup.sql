SELECT EXISTS (SELECT 1 FROM eft.lineup WHERE playerid = $1);
