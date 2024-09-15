package strategy

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/easy-football-tycoon/app"
)

type PostStrategyRequest struct {
	Formation            Formation            `json:"formation"`
	PlayingStyle         PlayingStyle         `json:"playing_style"`
	GameTempo            GameTempo            `json:"game_tempo"`
	PassingStyle         PassingStyle         `json:"passing_style"`
	DefensivePositioning DefensivePositioning `json:"defensive_position"`
	BuildUpPlay          BuildUpPlay          `json:"build_up_play"`
	AttackFocus          AttackFocus          `json:"attack_focus"`
	KeyPlayerUsage       KeyPlayerUsage       `json:"key_player_usage"`
}

type Formation string

const (
	FourFourTwo    Formation = "4-4-2"
	FourThreeThree Formation = "4-3-3"
	FourFiveOne    Formation = "4-5-1"
	FiveFourOne    Formation = "5-4-1"
	FiveThreeTwo   Formation = "5-3-2"
	ThreeFourThree Formation = "3-4-3"
	ThreeFiveTwo   Formation = "3-5-2"
)

type PlayingStyle string

const (
	Possession    PlayingStyle = "possession"
	CounterAttack PlayingStyle = "counter_attack"
	DirectPlay    PlayingStyle = "direct_play"
	HighPress     PlayingStyle = "high_press"
	LowBlock      PlayingStyle = "low_block"
)

type GameTempo string

const (
	FastTempo     GameTempo = "fast_tempo"
	BalancedTempo GameTempo = "balanced_tempo"
	SlowTempo     GameTempo = "slow_tempo"
)

type PassingStyle string

const (
	ShortPassing PassingStyle = "short"
	LongPassing  PassingStyle = "long"
)

type DefensivePositioning string

const (
	ZonalMarking DefensivePositioning = "zonal_marking"
	ManMarking   DefensivePositioning = "man_marking"
)

type BuildUpPlay string

const (
	PlayFromBack  BuildUpPlay = "play_from_back"
	LongClearance BuildUpPlay = "long_clearance"
)

type AttackFocus string

const (
	WidePlay    AttackFocus = "wide_play"
	CentralPlay AttackFocus = "central_play"
)

type KeyPlayerUsage string

const (
	ReferencePlayer KeyPlayerUsage = "reference_player"
	FreeRolePlayer  KeyPlayerUsage = "free_role_player"
)

func (h Handler) PostStrategy(c *gin.Context) {
	var req PostStrategyRequest
	if err := c.BindJSON(&req); err != nil {
		log.Printf("[PostStrategy] error parsing request: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	validFormations := map[Formation]bool{
		FourFourTwo:    true,
		FourThreeThree: true,
		FourFiveOne:    true,
		FiveFourOne:    true,
		FiveThreeTwo:   true,
		ThreeFourThree: true,
		ThreeFiveTwo:   true,
	}

	if !validFormations[req.Formation] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "formación inválida"})
		return
	}

	validPlayingStyles := map[PlayingStyle]bool{
		Possession:    true,
		CounterAttack: true,
		DirectPlay:    true,
		HighPress:     true,
		LowBlock:      true,
	}

	if !validPlayingStyles[req.PlayingStyle] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "estilo de juego inválido"})
		return
	}

	validGameTempos := map[GameTempo]bool{
		FastTempo:     true,
		BalancedTempo: true,
		SlowTempo:     true,
	}

	if !validGameTempos[req.GameTempo] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tempo del juego inválido"})
		return
	}

	validPassingStyles := map[PassingStyle]bool{
		ShortPassing: true,
		LongPassing:  true,
	}

	if !validPassingStyles[req.PassingStyle] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "estilo de pase inválido"})
		return
	}

	validDefensivePositioning := map[DefensivePositioning]bool{
		ZonalMarking: true,
		ManMarking:   true,
	}

	if !validDefensivePositioning[req.DefensivePositioning] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "posición defensiva inválida"})
		return
	}

	validBuildUpPlays := map[BuildUpPlay]bool{
		PlayFromBack:  true,
		LongClearance: true,
	}

	if !validBuildUpPlays[req.BuildUpPlay] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "construcción del juego inválida"})
		return
	}

	validAttackFocus := map[AttackFocus]bool{
		WidePlay:    true,
		CentralPlay: true,
	}

	if !validAttackFocus[req.AttackFocus] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "enfoque de ataque inválido"})
		return
	}

	validKeyPlayerUsage := map[KeyPlayerUsage]bool{
		ReferencePlayer: true,
		FreeRolePlayer:  true,
	}

	if !validKeyPlayerUsage[req.KeyPlayerUsage] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "uso del jugador clave inválido"})
		return
	}

	strategy := app.Strategy{
		Formation:            string(req.Formation),
		PlayingStyle:         string(req.PlayingStyle),
		GameTempo:            string(req.GameTempo),
		PassingStyle:         string(req.PassingStyle),
		DefensivePositioning: string(req.DefensivePositioning),
		BuildUpPlay:          string(req.BuildUpPlay),
		AttackFocus:          string(req.AttackFocus),
		KeyPlayerUsage:       string(req.KeyPlayerUsage),
	}

	err := h.app.PostStrategy(strategy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "strategy insertado correctamente"})
}
