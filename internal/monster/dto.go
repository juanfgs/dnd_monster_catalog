package monster

import (
	"github.com/juanfgs/dnd-monster-library/internal/armor"
	"github.com/juanfgs/dnd-monster-library/internal/proficiency"
	"github.com/juanfgs/dnd-monster-library/internal/speed"
	"github.com/juanfgs/dnd-monster-library/internal/stats"
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
	ChallengeRating float64 `json:"challenge_rating"`
	ProficiencyBonus int64 
	Proficiencies []proficiency.ProficiencyDTO
	Stats *stats.Stats
	Speed speed.SpeedDTO
	ArmorClass []armor.ArmorClassDTO `json:"armor_class"`
	XP int64
	// virtual attributes
	Strength int64 
	Dexterity int64 
	Constitution int64
	Intelligence int64
	Wisdom int64
	Charisma int64
}

func (d MonsterDTO) BuildModel() *Monster{
	proficiencies := proficiency.BuildModels(d.Proficiencies)
	armorClasses := armor.BuildModels(d.ArmorClass)
	speeds := speed.BuildModels(d.Speed)
	return &Monster{
		ID: d.ID,
		Index: d.Index,
		Name: d.Name, 
		Size: d.Size,
		Alignment: d.Alignment,
		ArmorClasses: armorClasses,
		HitPoints: d.HitPoints, 
		HitDice: d.HitDice, 
		HitPointsRoll: d.HitPointsRoll, 
		Languages: d.Languages,
		ChallengeRating: d.ChallengeRating, 
		ProficiencyBonus: d.ProficiencyBonus, 
		Proficiencies: proficiencies,
		Stats: &stats.Stats{
			Strength: d.Strength,
			Dexterity: d.Dexterity,
			Constitution: d.Constitution,
			Intelligence: d.Intelligence,
			Wisdom: d.Wisdom,
			Charisma: d.Charisma,
		},
		Speed: speeds,
		XP: d.XP, 
	}
}
