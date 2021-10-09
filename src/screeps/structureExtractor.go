package screeps

type StructureExtractor struct {
	OwnedStructure
}

func (extractor StructureExtractor) Cooldown() int {
	return extractor.ref.Get("cooldown").Int()
}
