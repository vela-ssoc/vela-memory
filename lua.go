package memory

import (
	"github.com/vela-ssoc/vela-kit/vela"
)

var (
	xEnv vela.Environment
)

func WithEnv(env vela.Environment) {
	xEnv = env
	sum := New()
	sum.Update()
	xEnv.Set("memory", sum)
}
