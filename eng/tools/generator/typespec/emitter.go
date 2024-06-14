package typespec

type TypeSpecEmitters string

const (
	TypeSpec_GO       TypeSpecEmitters = "@azure-tools/typespec-go"
	TypeSpec_AUTOREST TypeSpecEmitters = "@azure-tools/typespec-autorest"
	TypeSpec_CSHARP   TypeSpecEmitters = "@azure-tools/typespec-csharp"
	TypeSpec_PYTHON   TypeSpecEmitters = "@azure-tools/typespec-python"
	TypeSpec_TS       TypeSpecEmitters = "@azure-tools/typespec-ts"
)
