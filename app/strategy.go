package app

import "github.com/google/uuid"

type Strategy struct {
	StrategyId           uuid.UUID
	Formation            string
	PlayingStyle         string
	GameTempo            string
	PassingStyle         string
	DefensivePositioning string
	BuildUpPlay          string
	AttackFocus          string
	KeyPlayerUsage       string
}
