INSERT  into EFT.team_staff
    	(firstname, 
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
		rarity)
VALUES
   ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14);