SELECT playerid, appearances, blocks, saves, aerialduel, keypass, assists, chances, goals, mvp, rating
		FROM eft.player_stats
		WHERE playerid = $1;