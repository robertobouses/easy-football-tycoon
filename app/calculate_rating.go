package app

const (
	IncreaseRatingSlightly     = 0.11
	IncreaseRatingModerately   = 0.35
	IncreaseRatingConsiderably = 0.67
	IncreaseRatingDrastically  = 1.29

	DecreaseRatingSlightly     = -0.10
	DecreaseRatingModerately   = -0.32
	DecreaseRatingConsiderably = -0.64
	DecreaseRatingDrastically  = -1.36
)

// func IncreaseRatingSlightly() float64 {
// 	return +0.11

// }

// func IncreaseRatingModerately() float64 {
// 	return +0.35

// }
// func IncreaseRatingConsiderably(rating float64) float64 {
// 	return rating + 0.67

// }

// func IncreaseRatingDrastically(rating float64) float64 {
// 	return rating + 1.29

// }

// func DecreaseRatingSlightly(rating float64) float64 {
// 	return rating - 0.10

// }
// func DecreaseRatingModerately(rating float64) float64 {
// 	return rating - 0.32

// }
// func DecreaseRatingConsiderably(rating float64) float64 {
// 	return rating - 0.64

// }

// func DecreaseRatingDrastically(rating float64) float64 {
// 	return rating - 1.36

// }

// func (a AppService) IncreaseRatingSlightly(playerId uuid.UUID) (float64, error) {
// 	playerStats, err := a.statsRepo.GetPlayerStatsById(playerId)
// 	if err != nil {
// 		return 0, fmt.Errorf("error fetching player stats: %v", err)
// 	}

// 	newRating := playerStats.Rating + 0.11
// 	return newRating, nil
// }

// func (a AppService) IncreaseRatingModerately(playerId uuid.UUID) float64 {
// 	playerStats, err := a.statsRepo.GetPlayerStatsById(playerId)
// 	if err != nil {
// 		return 0
// 	}

// 	newRating := playerStats.Rating + 0.35
// 	return newRating
// }

// func (a AppService) IncreaseRatingConsiderably(playerId uuid.UUID) (float64, error) {
// 	playerStats, err := a.statsRepo.GetPlayerStatsById(playerId)
// 	if err != nil {
// 		return 0, fmt.Errorf("error fetching player stats: %v", err)
// 	}

// 	newRating := playerStats.Rating + 0.67
// 	return newRating, nil
// }

// func (a AppService) IncreaseRatingDrastically(playerId uuid.UUID) (float64, error) {
// 	playerStats, err := a.statsRepo.GetPlayerStatsById(playerId)
// 	if err != nil {
// 		return 0, fmt.Errorf("error fetching player stats: %v", err)
// 	}

// 	newRating := playerStats.Rating + 1.29
// 	return newRating, nil
// }

// func (a AppService) DecreaseRatingSlightly(playerId uuid.UUID) (float64, error) {
// 	playerStats, err := a.statsRepo.GetPlayerStatsById(playerId)
// 	if err != nil {
// 		return 0, fmt.Errorf("error fetching player stats: %v", err)
// 	}

// 	newRating := playerStats.Rating - 0.10
// 	return newRating, nil
// }

// func (a AppService) DecreaseRatingModerately(playerId uuid.UUID) (float64, error) {
// 	playerStats, err := a.statsRepo.GetPlayerStatsById(playerId)
// 	if err != nil {
// 		return 0, fmt.Errorf("error fetching player stats: %v", err)
// 	}

// 	newRating := playerStats.Rating - 0.32
// 	return newRating, nil
// }

// func (a AppService) DecreaseRatingConsiderably(playerId uuid.UUID) (float64, error) {
// 	playerStats, err := a.statsRepo.GetPlayerStatsById(playerId)
// 	if err != nil {
// 		return 0, fmt.Errorf("error fetching player stats: %v", err)
// 	}

// 	newRating := playerStats.Rating - 0.64
// 	return newRating, nil
// }

// func (a AppService) DecreaseRatingDrastically(playerId uuid.UUID) (float64, error) {
// 	playerStats, err := a.statsRepo.GetPlayerStatsById(playerId)
// 	if err != nil {
// 		return 0, fmt.Errorf("error fetching player stats: %v", err)
// 	}

// 	newRating := playerStats.Rating - 1.36
// 	return newRating, nil
// }
