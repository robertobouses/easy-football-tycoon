SELECT
			prospectid,
			prospectname,
			position,
			age,
			fee,
			salary,
			technique,
			mental,
			physique,
			injurydays,
			job,
			rarity
		FROM
			eft.prospect
		WHERE
			$1 > rarity
		ORDER BY
			RANDOM()
		LIMIT 1;