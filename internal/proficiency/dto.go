package proficiency

type ProficiencyDTOMetadata struct {
	Index string
	Name string
	Url string
}

type ProficiencyDTO struct {
	Value int64 
	Metadata ProficiencyDTOMetadata `json:"proficiency"` 
}

func BuildModels(p []ProficiencyDTO) []*Proficiency {
	proficiencies := make([]*Proficiency, 0)
	for _, e := range(p) {
		proficiencies = append(proficiencies, e.BuildModel())
	}
	return proficiencies
}

func (p ProficiencyDTO) BuildModel() *Proficiency {
	return &Proficiency{
		Name: p.Metadata.Name,
		Value: p.Value,
	}
}
