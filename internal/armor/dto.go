package armor 

type ArmorClassDTO struct {
	Type string 
	Value int64
}

func BuildModels(p []ArmorClassDTO) []ArmorClass {
	armorClasses := make([]ArmorClass, 0)
	for _, e := range(p) {
		armorClasses = append(armorClasses, e.BuildModel())
	}
	return armorClasses
}

func (p ArmorClassDTO) BuildModel() ArmorClass {
	return ArmorClass{
		Type: p.Type,
	}
}
