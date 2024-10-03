INSERT  into EFT.team
    	(firstname,
		lastname,
		nationality,
		position,
		age,
		fee,
		salary,
		technique,
		mental,
		physique,
		injurydays,
		lined,
		familiarity,
		fitness,
		happiness)
VALUES
   ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15);