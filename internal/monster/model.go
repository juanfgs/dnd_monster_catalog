package monster
import(
	"github.com/juanfgs/dnd-monster-library/internal/stats"
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

type Proficiency struct {
	Value int64 
	Name string
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
	Stats *stats.Stats
	XP int64

	// virtual attributes
	strength int64 
	dexterity int64 
	constitution int64
	intelligence int64
	wisdom int64
	charisma int64

}


func (m *Monster) LoadStats() {
	m.Stats = &stats.Stats{
		Strength: m.strength,
		Dexterity: m.dexterity,
		Constitution: m.constitution,
		Intelligence: m.intelligence,
		Wisdom: m.wisdom,
		Charisma: m.charisma,
	}
}
