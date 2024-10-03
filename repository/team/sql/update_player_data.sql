UPDATE eft.team
		SET 
			firstname = COALESCE($1, firstname),
			lastname = COALESCE($1, lastname),
			nationality = COALESCE($1, nationality),
			position = COALESCE($2, position),
			age = COALESCE($3, age),
			fee = COALESCE($4, fee),
			salary = COALESCE($5, salary),
			technique = COALESCE($6, technique),
			mental = COALESCE($7, mental),
			physique = COALESCE($8, physique),
		    injurydays = COALESCE($9, 0) + COALESCE(injurydays, 0),
			lined = COALESCE($10, lined),
			familiarity = GREATEST(COALESCE(familiarity, 0) - COALESCE($11, 0), 0),
			fitness = LEAST(GREATEST(COALESCE($12, fitness), 0), COALESCE(fitness, 0) - COALESCE($12, 0)),
			happiness = GREATEST(COALESCE(happiness, 0) - COALESCE($13, 0), 0)

		WHERE playerid = $14;


		