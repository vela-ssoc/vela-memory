package memory

import (
	"github.com/vela-ssoc/vela-kit/vela"
)

var (
	xEnv vela.Environment
	_G   *summary
)

func WithEnv(env vela.Environment) {
	xEnv = env
	define(env.R())
	sum := New()
	sum.Update()
	xEnv.Set("memory", sum)
	_G = sum
}
