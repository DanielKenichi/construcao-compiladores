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
	INVALIDO
)

type SymbolTable struct {
	Table map[string]*Symbol
}

type Symbol struct {
	SymbolType Type
	//InnerTable é nil se o simbolo nao for REGISTRO, FUNCAO ou PROCEDIMENTO
	InnerTable *SymbolTable

	//Só deve ser usado se o símbolo for uma FUNCAO ou PROCEDIMENTO
	//Salvamos como string pois se o retorno for um registro precisamos
	//do identificador desse registro
	ReturnType string
}

func New() *SymbolTable {
	return &SymbolTable{
		Table: make(map[string]*Symbol),
	}
}

func (s *SymbolTable) NewSymbol(name string, symbolType Type) *Symbol {

	var innerTable *SymbolTable = nil

	if symbolType == REGISTRO || symbolType == REGISTRO_VAR ||
		symbolType == FUNCAO || symbolType == PROCEDIMENTO {

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

func (s *SymbolTable) AddReturnTypeToSymbol(name string, returnType string) {
	symbol := s.GetSymbol(name)

	symbol.ReturnType = returnType
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
