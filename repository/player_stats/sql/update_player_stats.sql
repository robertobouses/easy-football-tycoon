INSERT INTO eft.player_stats (player_id, appearances, chances, assists, goals, mvp, rating)
VALUES ($7, $1, $2, $3, $4, $5, $6)
ON CONFLICT (player_id)
DO UPDATE SET 
    appearances = eft.player_stats.appearances + EXCLUDED.appearances,
    chances = eft.player_stats.chances + EXCLUDED.chances,
    assists = eft.player_stats.assists + EXCLUDED.assists,
    goals = eft.player_stats.goals + EXCLUDED.goals,
    mvp = eft.player_stats.mvp + EXCLUDED.mvp,
    rating = (eft.player_stats.rating * eft.player_stats.appearances + EXCLUDED.rating) / (eft.player_stats.appearances + EXCLUDED.appearances);
