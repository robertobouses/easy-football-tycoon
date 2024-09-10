SELECT
			signingsid,
			signingsname,
			position,
			age,
			fee,
			salary,
			technique,
			mental,
			physique,
			injurydays,
			rarity
		FROM
			eft.signings
		WHERE
			$1 > rarity
		ORDER BY
			RANDOM()
		LIMIT 1;