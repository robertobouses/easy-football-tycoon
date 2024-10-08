SELECT
			staffid,
			firstname,
			lastname,
			nationality,
			job,
			age,
			fee,
			salary,
			training,
			finances,
			scouting,
			physiotherapy,
			knowledge,
			intelligence,
			rarity
		FROM
			eft.staff
		WHERE
			$1 > rarity
		ORDER BY
			RANDOM()
		LIMIT 1;