package strategy

import (
	//"database/sql"
	"log"

	"github.com/robertobouses/easy-football-tycoon/app"
)

func (r *repository) GetStrategy() (app.Strategy, error) {

	row := r.getStrategy.QueryRow()
	var strategy app.Strategy
	if err := row.Scan(
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
		log.Printf("Error al escanear las filas GetStrategy: %v", err)
		return app.Strategy{}, err
	}

	return strategy, nil
}
