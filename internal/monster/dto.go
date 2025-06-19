package monster

import(
	"github.com/juanfgs/dnd-monster-library/internal/stats"
	"github.com/juanfgs/dnd-monster-library/internal/proficiency"
)

type MonsterDTO struct {
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
	Proficiencies []proficiency.ProficiencyDTO
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

func (d MonsterDTO) BuildModel() *Monster{
	proficiencies := proficiency.BuildModels(d.Proficiencies)
	return &Monster{
		ID: d.ID,
		Index: d.Index,
		Name: d.Name, 
		Size: d.Size,
		Alignment: d.Alignment,
		HitPoints: d.HitPoints, 
		HitDice: d.HitDice, 
		HitPointsRoll: d.HitPointsRoll, 
		Languages: d.Languages,
		ProficiencyBonus: d.ProficiencyBonus, 
		Proficiencies: proficiencies,
		Stats: &stats.Stats{
		Strength: d.strength,
		Dexterity: d.dexterity,
		Constitution: d.constitution,
		Intelligence: d.intelligence,
		Wisdom: d.wisdom,
		Charisma: d.charisma,
		},
		XP: d.XP, 
	}
}
