package monster
import(
	"github.com/juanfgs/dnd-monster-library/internal/stats"
	"github.com/juanfgs/dnd-monster-library/internal/proficiency"
)

type Speed struct {
	Walk string
	Swim string
	Fly string
	Climb string
}

type ArmorClass struct {
	Type string
	Value int64
}

type Monster struct {
	ID string
	Index string
	Name string
	Size string
	Alignment string
	HitPoints int64
	HitDice string
	HitPointsRoll string
	Languages string
	ProficiencyBonus int64
	Proficiencies []proficiency.Proficiency
	Stats *stats.Stats
	XP int64
}

