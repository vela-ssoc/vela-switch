package vswitch

import (
	"github.com/vela-ssoc/vela-kit/vela"
	"github.com/vela-ssoc/vela-kit/lua"
)

var xEnv vela.Environment

/*
	local switch = vela.switch("name")
	local s = switch._{
		["baidu.com"] = pipe,
	}

	s.once(true)
	s.case("name eq baidu.com" , "icmp > 188").pipe()
	s.case("name eq baidu.com" , "icmp > 188").pipe()
	s.case("name eq baidu.com" , "icmp > 188").pipe()
	s.case("name eq baidu.com" , "icmp > 188").pipe()
	s.case("name eq baidu.com" , "icmp > 188").pipe()
	s.case("name eq baidu.com" , "icmp > 188").pipe()
	s.default(pipe1 , pipe2 , pipe3)
	s.match(app)
*/

func newLuaSwitchL(L *lua.LState) int {
	s := NewSwitchL(L)
	L.Push(s)
	return 1
}

func WithEnv(env vela.Environment) {
	xEnv = env
	xEnv.Set("switch", lua.NewFunction(newLuaSwitchL))
}
