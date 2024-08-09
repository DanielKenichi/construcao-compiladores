package symboltable

type Type int32

const (
	INTEIRO Type = iota
	REAL
	LITERAL
	LOGICO
	REGISTRO     //O tipo do registro em si
	REGISTRO_VAR //declaracoes de um tipo de um registro
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

	//Só deve ser usado se o símbolo for uma FUNCAO
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
	return s.Table[name].SymbolType
}

func (s *SymbolTable) GetSymbol(name string) *Symbol {
	return s.Table[name]
}
