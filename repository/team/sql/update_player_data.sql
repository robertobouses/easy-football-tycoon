UPDATE eft.team
		SET 
			playername = COALESCE($1, playername),
			position = COALESCE($2, position),
			age = COALESCE($3, age),
			fee = COALESCE($4, fee),
			salary = COALESCE($5, salary),
			technique = COALESCE($6, technique),
			mental = COALESCE($7, mental),
			physique = COALESCE($8, physique),
		    injurydays = COALESCE($9, 0) + COALESCE(injurydays, 0),
			lined = COALESCE($10, lined)
		WHERE playerid = $11;