package vocabulary

import "strconv"

/*Implementacao propria de um vocabulario em GO
  baseado na implementacao existente em java*/

type Vocabulary interface {
	GetDisplayName(tokenType int) *string
	GetSymbolicName(tokenType int) *string
	GetLiteralName(tokenType int) *string
}

type vocabulary struct {
	LiteralNames  []string
	SymbolicNames []string
	DisplayNames  []string
}

func New(literalNames []string, symbolicNames []string) Vocabulary {
	return &vocabulary{
		LiteralNames:  literalNames,
		SymbolicNames: symbolicNames,
		DisplayNames:  []string{},
	}
}

func (v *vocabulary) GetDisplayName(tokenType int) *string {
	var displayName *string

	displayName = v.GetLiteralName(tokenType)

	if displayName != nil {
		return displayName
	}

	displayName = v.GetSymbolicName(tokenType)

	if displayName != nil {
		return displayName
	} else {
		stringToken := strconv.Itoa(tokenType)
		return &stringToken
	}
}

func (v *vocabulary) GetSymbolicName(tokenType int) *string {

	if tokenType >= 0 && tokenType < len(v.SymbolicNames) {
		return &v.SymbolicNames[tokenType]
	} else if tokenType == -1 {
		eof := "EOF"
		return &eof
	} else {
		return nil
	}
}

func (v *vocabulary) GetLiteralName(tokenType int) *string {

	if tokenType >= 0 && tokenType < len(v.LiteralNames) {
		return &v.LiteralNames[tokenType]
	} else {
		return nil
	}
}
