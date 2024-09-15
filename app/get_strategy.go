package app

import "log"

func (a AppService) GetStrategy() (Strategy, error) {
	strategy, err := a.strategyRepo.GetStrategy()
	if err != nil {
		log.Println("Error al extraer GetStrategy", err)
		return Strategy{}, err
	}
	return strategy, nil
}
