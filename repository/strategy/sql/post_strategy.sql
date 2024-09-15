INSERT INTO eft.strategy (
    formation, 
    playing_style, 
    game_tempo, 
    passing_style, 
    defensive_positioning, 
    build_up_play, 
    attack_focus, 
    key_player_usage
) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8);
