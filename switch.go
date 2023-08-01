package vswitch

import (
	"fmt"
	cond "github.com/vela-ssoc/vela-cond"
	"github.com/vela-ssoc/vela-kit/lua"
	"github.com/vela-ssoc/vela-kit/pipe"
)

type Switch struct {
	co      *lua.LState
	Field   string
	Method  string
	Before  []*cond.Cond
	Debug   bool
	Break   bool
	Cases   []*Case
	Default *pipe.Chains
}

func (s *Switch) Append(c *Case) {
	s.Cases = append(s.Cases, c)
}

func (s *Switch) Len() int {
	return len(s.Cases)
}

func (s *Switch) ByIgnoreAndCallback(v interface{}, ignore *cond.Cond, fn func(cnd string, id uint8)) {
	n := s.Len()
	if n == 0 {
		return
	}

	for i := 0; i < n; i++ {
		ca := s.Cases[i]
		cnd := ca.cnd.String()
		if ignore != nil {
			if ignore.Match(cnd) {
				continue
			}
		}

		if !ca.Match(i+1, v, s.co) {
			continue
		}

		fn(cnd, uint8(i+1))
		if ca.over {
			return
		}
	}
}

func (s *Switch) Prefix() string {
	if s.Field == "" {
		return ""
	}

	method := "="
	if s.Method != "" {
		method = s.Method
	}
	return fmt.Sprintf("%s %s ", s.Field, method)
}

func (s *Switch) before(v interface{}) bool {
	n := len(s.Before)
	if n == 0 {
		return true
	}

	for i := 0; i < n; i++ {
		if s.Before[i].Match(v) {
			return true
		}
	}

	return false
}

func (s *Switch) PrepareDebug() func(int, string) {
	if !s.Debug {
		return func(i int, info string) {}
	}

	return func(i int, info string) {
		xEnv.Debugf("switch.%d %s", i, info)
	}
}

func (s *Switch) Do(v interface{}) {
	flag := false
	debug := s.PrepareDebug()

	n := s.Len()
	if n == 0 {
		goto done
	}

	if !s.before(v) {
		goto done
	}

	for i := 0; i < n; i++ {
		ca := s.Cases[i]
		if !ca.Match(i+1, v, s.co) {
			debug(i, "not match")
			continue
		}

		flag = true
		debug(i, "match")
		if ca.over {
			return
		}
	}

done:
	if !flag && s.Default != nil {
		if err := s.Default.Case(v, 0, "*", s.co); err != nil {
			xEnv.Errorf("vela.switch.default todo fail %v", err)
		}
	}

}
