package strategy

import (
	"log"

	"github.com/robertobouses/easy-football-tycoon/app"
)

func (r *repository) PostStrategy(strategy app.Strategy) error {
	_, err := r.postStrategy.Exec(
		strategy.Formation,
		strategy.PlayingStyle,
		strategy.GameTempo,
		strategy.PassingStyle,
		strategy.DefensivePositioning,
		strategy.BuildUpPlay,
		strategy.AttackFocus,
		strategy.KeyPlayerUsage,
	)

	if err != nil {
		log.Print("Error en PostStrategy repo", err)
		return err
	}
	log.Println("Después de ejecutar la consulta preparada")
	return nil
}
