SELECT
			prospect_id,
			prospect_name,
			position,
			age,
			fee,
			salary,
			technique,
			mental,
			physique,
			injury_days,
			job,
			rarity
		FROM
			prospects
		WHERE
			$1 > rarity
		ORDER BY
			RANDOM()
		LIMIT 1;