package vswitch

import (
	"fmt"
	cond "github.com/vela-ssoc/vela-cond"
	"github.com/vela-ssoc/vela-kit/pipe"
	"github.com/vela-ssoc/vela-kit/lua"
)

func (s *Switch) String() string                         { return fmt.Sprintf("vela.switch:%p", s) }
func (s *Switch) Type() lua.LValueType                   { return lua.LTObject }
func (s *Switch) AssertFloat64() (float64, bool)         { return 0, false }
func (s *Switch) AssertString() (string, bool)           { return "", false }
func (s *Switch) AssertFunction() (*lua.LFunction, bool) { return nil, false }
func (s *Switch) Peek() lua.LValue                       { return s }

func (s *Switch) NewSwitchByLTab(tab *lua.LTable) {
	prefix := s.Prefix()
	tab.Range(func(key string, val lua.LValue) {
		cnd := cond.New(prefix + key)
		px := pipe.New()
		px.LValue(val)
		s.Append(&Case{over: false, cnd: cnd, todo: px})
	})
}

func (s *Switch) caseL(L *lua.LState) int {
	c := NewCaseL(L)
	s.Append(c)
	L.Push(c)
	return 1
}

func (s *Switch) doL(L *lua.LState) int {
	n := L.GetTop()
	if n == 0 {
		return 0
	}

	for i := 1; i <= n; i++ {
		s.Do(L.Get(i))
	}

	return 0
}

func (s *Switch) breakL(L *lua.LState) int {
	s.Break = L.IsTrue(1)
	return 0
}

func (s *Switch) debugL(L *lua.LState) int {
	s.Debug = L.IsTrue(1)
	return 0
}

func (s *Switch) initL(L *lua.LState) int {
	tab := L.CheckTable(1)
	s.NewSwitchByLTab(tab)
	return 0
}

func (s *Switch) beforeL(L *lua.LState) int {
	cnd := cond.CheckMany(L)
	s.Before = append(s.Before, cnd)
	return 0
}

func (s *Switch) defaultL(L *lua.LState) int {
	s.Default = pipe.NewByLua(L, pipe.Env(xEnv))
	return 0
}

func (s *Switch) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "_":
		return lua.NewFunction(s.initL)
	case "case":
		return lua.NewFunction(s.caseL)
	case "match":
		return lua.NewFunction(s.doL)
	case "before":
		return lua.NewFunction(s.beforeL)
	case "once":
		return lua.NewFunction(s.breakL)
	case "debug":
		return lua.NewFunction(s.debugL)
	case "default":
		return lua.NewFunction(s.defaultL)
	}

	return lua.LNil
}

func (s *Switch) MetaTable(L *lua.LState, key lua.LValue) lua.LValue {
	switch key.Type() {
	case lua.LTNil:
		return lua.NewFunction(s.initL)
	case lua.LTString:
		return s.Index(L, key.String())
	case lua.LTNumber:
		idx := int(key.(lua.LNumber)) - 1
		if idx < 0 || idx >= s.Len() {
			return s.Cases[idx]
		}
	case lua.LTInt:
		idx := int(key.(lua.LInt)) - 1
		if idx < 0 || idx >= s.Len() {
			return s.Cases[idx]
		}
	}

	return lua.LNil
}

func NewL(L *lua.LState) *Switch {
	return &Switch{co: xEnv.Clone(L)}
}

func NewSwitchL(L *lua.LState) *Switch {
	field := L.IsString(1)
	method := L.IsString(2)

	s := &Switch{
		co:     xEnv.Clone(L),
		Field:  field,
		Method: method,
	}

	return s
}
