UPDATE eft.team
		SET 
			firstname = COALESCE($1, firstname),
			lastname = COALESCE($2, lastname),
			nationality = COALESCE($3, nationality),
			position = COALESCE($4, position),
			age = COALESCE($5, age),
			fee = COALESCE($6, fee),
			salary = COALESCE($7, salary),
			technique = GREATEST(0, LEAST(COALESCE(technique, 0) + COALESCE($8, 0),100)),
			mental = GREATEST(0, LEAST(100, COALESCE(mental, 0) + COALESCE($9, 0))),
			physique = GREATEST(0, LEAST(COALESCE(physique, 0) + COALESCE($10, 0),100)),
		   	injurydays = GREATEST(COALESCE(injurydays, 0) + COALESCE($11, 0), 0),
			lined = COALESCE($12, lined),
			familiarity = LEAST(GREATEST(COALESCE(familiarity, 0) - COALESCE($13, 0), 0),100),
			fitness=  LEAST(100, GREATEST(0, COALESCE(fitness, 0) - COALESCE($14, 0))),
			happiness = LEAST(GREATEST(COALESCE(happiness, 0) - COALESCE($15, 0), 0), 100)

		WHERE playerid = $16;


		