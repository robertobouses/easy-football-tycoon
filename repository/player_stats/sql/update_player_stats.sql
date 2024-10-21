INSERT INTO eft.player_stats (playerid, appearances, blocks, saves, aerialduel, keypass, assists, chances, goals, mvp, rating)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
ON CONFLICT (playerid) DO UPDATE SET 
    appearances = player_stats.appearances + COALESCE(EXCLUDED.appearances, 0),
    blocks = player_stats.blocks + COALESCE(EXCLUDED.blocks, 0),
    saves = player_stats.saves + COALESCE(EXCLUDED.saves, 0),
    aerialduel = player_stats.aerialduel + COALESCE(EXCLUDED.aerialduel, 0),
    keypass = player_stats.keypass + COALESCE(EXCLUDED.keypass, 0),
    assists = player_stats.assists + COALESCE(EXCLUDED.assists, 0),
    chances = player_stats.chances + COALESCE(EXCLUDED.chances, 0),
    goals = player_stats.goals + COALESCE(EXCLUDED.goals, 0),
    mvp = player_stats.mvp + COALESCE(EXCLUDED.mvp, 0),
    rating = CASE 
                WHEN (player_stats.appearances + COALESCE(EXCLUDED.appearances, 0)) > 0 
                THEN (COALESCE(player_stats.rating, 0) * player_stats.appearances + COALESCE(EXCLUDED.rating, 0)) / 
                     (player_stats.appearances + COALESCE(EXCLUDED.appearances, 0))
                ELSE 0 
             END;