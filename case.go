package vswitch

import (
	cond "github.com/vela-ssoc/vela-cond"
	"github.com/vela-ssoc/vela-kit/lua"
	"github.com/vela-ssoc/vela-kit/pipe"
)

type Case struct {
	over bool
	cnd  *cond.Cond
	todo *pipe.Chains
}

func (c *Case) Field(key string) string {
	switch key {
	case "raw":
		return c.cnd.String()
	}
	return ""
}

func (c *Case) Match(idx int, v interface{}, co *lua.LState) bool {
	if c.cnd == nil {
		return false
	}

	if !c.cnd.Match(v) {
		return false
	}

	if c.todo == nil {
		return true
	}

	if err := c.todo.Case(v, idx, c.cnd.String(), co); err != nil {
		xEnv.Errorf("vela.select.todo fail %v", err)
	}

	return true
}

func NewCase() *Case {
	return &Case{}
}
