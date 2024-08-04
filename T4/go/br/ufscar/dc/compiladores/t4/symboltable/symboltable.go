package symboltable

import "log"

type Type int32

const (
	INTEIRO Type = iota
	REAL
	LITERAL
	LOGICO
	REGISTRO
	REGISTRO_VAR
	PONTEIRO
	ENDERECO
	FUNCAO
	PROCEDIMENTO
	TIPO
	INVALIDO
)

type SymbolTable struct {
	Table map[string]*Symbol
}

type Symbol struct {
	SymbolType Type
	//InnerTable is nil if Symbol is not a REGISTRO
	InnerTable *SymbolTable
}

func New() *SymbolTable {
	return &SymbolTable{
		Table: make(map[string]*Symbol),
	}
}

func (s *SymbolTable) NewSymbol(name string, symbolType Type) *Symbol {

	var innerTable *SymbolTable = nil

	if symbolType == REGISTRO || symbolType == REGISTRO_VAR {
		innerTable = New()
	}

	return &Symbol{
		SymbolType: symbolType,
		InnerTable: innerTable,
	}
}

func (s *SymbolTable) AddSymbol(name string, symbolType Type) {

	log.Printf("Adding symbol %v with type %v", name, symbolType)

	symbol := s.NewSymbol(name, symbolType)

	s.Table[name] = symbol
}

func (s *SymbolTable) AddInnerTableToRegVar(name string, table SymbolTable) {
	symbol := s.GetSymbol(name)

	symbol.InnerTable = &table
}

func (s *SymbolTable) Exists(name string) bool {
	_, ok := s.Table[name]

	return ok
}

func (s *SymbolTable) GetType(name string) Type {
	log.Printf("Getting type for %v", name)
	return s.Table[name].SymbolType
}

func (s *SymbolTable) GetSymbol(name string) *Symbol {
	return s.Table[name]
}
