INSERT  into EFT.team
    	(playerName,
		position,
		age,
		fee,
		salary,
		technique,
		mental,
		physique,
		injuryDays,
		lined)
VALUES
   ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);
    