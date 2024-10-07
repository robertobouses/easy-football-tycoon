UPDATE eft.team
		SET 
			firstname = COALESCE($1, firstname),
			lastname = COALESCE($2, lastname),
			nationality = COALESCE($3, nationality),
			position = COALESCE($4, position),
			age = COALESCE($5, age),
			fee = COALESCE($6, fee),
			salary = COALESCE($7, salary),
			technique = COALESCE($8, technique),
			mental = COALESCE($9, mental),
			physique = COALESCE($10, physique),
		    injurydays = COALESCE($11, 0) + COALESCE(injurydays, 0),
			lined = COALESCE($12, lined),
			familiarity = GREATEST(COALESCE(familiarity, 0) - COALESCE($13, 0), 0),
			fitness = LEAST(GREATEST(COALESCE($14, fitness), 0), COALESCE(fitness, 0) - COALESCE($14, 0)),
			happiness = GREATEST(COALESCE(happiness, 0) - COALESCE($15, 0), 0)

		WHERE playerid = $16;


		