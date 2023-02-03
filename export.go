package vswitch

import "github.com/vela-ssoc/vela-kit/lua"

func CheckSwitch(L *lua.LState, idx int) *Switch {
	v := L.Get(idx).Peek()

	if v.Type() != lua.LTObject {
		L.RaiseError("%s must be switch object , but got %s", idx, v.Type().String())
		return nil
	}

	if s, ok := v.(*Switch); ok {
		return s
	}

	L.RaiseError("%d invalid switch object", idx)
	return nil

}

func IsSwitch(val lua.LValue) *Switch {
	if val.Type() != lua.LTObject {
		return nil
	}

	if s, ok := val.(*Switch); ok {
		return s
	}
	return nil
}
