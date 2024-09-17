CREATE TABLE eft.strategy (
    strategy_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    formation VARCHAR(50) NOT NULL,
    playing_style VARCHAR(50) NOT NULL,
    game_tempo VARCHAR(50) NOT NULL,
    passing_style VARCHAR(50) NOT NULL,
    defensive_positioning VARCHAR(50) NOT NULL,
    build_up_play VARCHAR(50) NOT NULL,
    attack_focus VARCHAR(50) NOT NULL,
    key_player_usage VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
