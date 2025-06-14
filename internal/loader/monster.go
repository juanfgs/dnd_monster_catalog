package loader

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

type Stats struct {
	Strength int64 
	Dexterity int64 
	Constitution int64
	Intelligence int64
	Wisdom int64
	Charisma int64
}

type Monster struct {
	Index string
	Name string
	Size string
	Alignment string
	ArmorClass []ArmorClass
	HitPoints int64
	HitDice string
	HitPointsRoll string
	Speed Speed
	Stats *Stats
	Proficiencies []Proficiency
	Languages string
	ProficiencyBonus int64
	XP int64
}


