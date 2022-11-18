package enumfilter

type EnumRemove string

const (
	EnumRemoveA EnumRemove = "A"
	EnumRemoveB EnumRemove = "B"
	EnumRemoveC EnumRemove = "C"
)

func PossibleEnumRemoveValues() []EnumRemove {
	return []EnumRemove{
		EnumRemoveA,
		EnumRemoveB,
		EnumRemoveC,
	}
}
