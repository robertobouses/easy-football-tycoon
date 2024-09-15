package strategy

import (
	//"database/sql"
	"log"

	"github.com/robertobouses/easy-football-tycoon/app"
)

func (r *repository) GetStrategy() (app.Strategy, error) {

	rows, err := r.getStrategy.Query()
	if err != nil {
		return app.Strategy{}, err
	}
	defer rows.Close()

	var strategy app.Strategy
	if err := rows.Scan(
		&strategy.StrategyId,
		&strategy.Formation,
		&strategy.PlayingStyle,
		&strategy.GameTempo,
		&strategy.PassingStyle,
		&strategy.DefensivePositioning,
		&strategy.BuildUpPlay,
		&strategy.AttackFocus,
		&strategy.KeyPlayerUsage,
	); err != nil {
		log.Printf("Error al escanear las filas: %v", err)
		return app.Strategy{}, err
	}

	return strategy, nil
}
