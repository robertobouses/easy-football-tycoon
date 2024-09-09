SELECT
			staff,
			staffname,
			job,
			age,
			fee,
			salary,
			training,
			finances,
			scouting,
			physiotherapy,
			rarity
		FROM
			eft.staff
		WHERE
			$1 > rarity
		ORDER BY
			RANDOM()
		LIMIT 1;