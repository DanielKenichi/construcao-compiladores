package scope

import "github.com/DanielKenichi/construcao-compiladores/T3/go/br/ufscar/dc/compiladores/t3/symboltable"

type Scope struct {
	Stack []*symboltable.SymbolTable
}

func New() *Scope {
	return &Scope{
		Stack: make([]*symboltable.SymbolTable, 0),
	}
}

func (s *Scope) Push(table *symboltable.SymbolTable) {
	s.Stack = append(s.Stack, table)
}

func (s *Scope) Pop() *symboltable.SymbolTable {
	table := s.Stack[len(s.Stack)-1]

	s.Stack = s.Stack[:len(s.Stack)-1]

	return table
}

func (s *Scope) NewScope() {
	s.Push(symboltable.New())
}

func (s *Scope) RemoveScope() *symboltable.SymbolTable {
	return s.Pop()
}

func (s *Scope) CurrentScope() *symboltable.SymbolTable {
	return s.Stack[len(s.Stack)-1]
}
