package proficiency

import (
	"strings"
)

type ProficiencyDTOMetadata struct {
	Index string
	Name string
	Url string
}

type ProficiencyDTO struct {
	Value int64 
	Metadata ProficiencyDTOMetadata `json:"proficiency"` 
}

func BuildModels(p []ProficiencyDTO) []Proficiency {
	proficiencies := make([]Proficiency, 0)
	for _, e := range(p) {
		proficiencies = append(proficiencies, e.BuildModel())
	}
	return proficiencies
}

func (p ProficiencyDTO) BuildModel() Proficiency {
	typeAndNamePair := strings.Split(p.Metadata.Name, ": ")
	
	return Proficiency{
		Type: typeAndNamePair[0],
		Name: typeAndNamePair[1],
	}
}
