package monster

import (
	"github.com/juanfgs/dnd-monster-library/internal/armor"
	"github.com/juanfgs/dnd-monster-library/internal/proficiency"
	"github.com/juanfgs/dnd-monster-library/internal/speed"
	"github.com/juanfgs/dnd-monster-library/internal/stats"
)

type Monster struct {
	ID string
	Index string
	Name string
	Size string
	Alignment string
	ArmorClasses []armor.ArmorClass
	HitPoints int64
	HitDice string
	HitPointsRoll string
	Languages string
	ChallengeRating float64 
	ProficiencyBonus int64
	Proficiencies []proficiency.Proficiency
	Speed []speed.Speed
	Stats *stats.Stats
	XP int64
}

