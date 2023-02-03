package vswitch

import (
	"fmt"
	cond "github.com/vela-ssoc/vela-cond"
	"github.com/vela-ssoc/vela-kit/lua"
	"github.com/vela-ssoc/vela-kit/pipe"
)

func (c *Case) String() string                         { return fmt.Sprintf("vela.Case:%p", c) }
func (c *Case) Type() lua.LValueType                   { return lua.LTObject }
func (c *Case) AssertFloat64() (float64, bool)         { return 0, false }
func (c *Case) AssertString() (string, bool)           { return "", false }
func (c *Case) AssertFunction() (*lua.LFunction, bool) { return nil, false }
func (c *Case) Peek() lua.LValue                       { return c }

func (c *Case) pipeL(L *lua.LState) int {
	pip := pipe.NewByLua(L)
	c.todo = pip
	L.Push(c)
	return 1
}

func (c *Case) breakL(L *lua.LState) int {
	c.over = L.IsTrue(1)
	L.Push(c)
	return 1
}

func (c *Case) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "pipe":
		return lua.NewFunction(c.pipeL)
	case "over":
		return lua.NewFunction(c.breakL)

	}
	return lua.LNil
}

func NewCaseL(L *lua.LState) *Case {
	cnd := cond.CheckMany(L)
	return &Case{over: false, cnd: cnd}
}
