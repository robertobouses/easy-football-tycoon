SELECT player_id, appearances, blocks, saves, aerialduel, keypass, assists, chances, goals, mvp, rating
		FROM eft.player_stats
		WHERE player_id = $1;