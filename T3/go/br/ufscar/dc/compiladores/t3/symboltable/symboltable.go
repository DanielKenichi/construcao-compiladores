package symboltable

type Type int32

const (
	INTEIRO Type = iota
	REAL
	LITERAL
	LOGICO
	REGISTRO
	PONTEIRO
	FUNCAO
	PROCEDIMENTO
	INVALIDO
)

type SymbolTable struct {
	Table map[string]Type
}

func New() *SymbolTable {
	return &SymbolTable{
		Table: make(map[string]Type),
	}
}

func (s *SymbolTable) AddSymbol(name string, symbolType Type) {
	s.Table[name] = symbolType
}

func (s *SymbolTable) Exists(name string) bool {
	_, ok := s.Table[name]

	return ok
}

func (s *SymbolTable) GetType(name string) Type {
	return s.Table[name]
}
