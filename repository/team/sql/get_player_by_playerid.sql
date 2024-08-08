SELECT playerid, playername, position, technique, mental, physique
		FROM eft.team
		WHERE playerid = $1;