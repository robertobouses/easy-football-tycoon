SELECT playerid, firstname, lastname, nationality, position, technique, mental, physique
		FROM eft.team
		WHERE playerid = $1;