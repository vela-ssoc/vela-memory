package memory

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"github.com/vela-ssoc/vela-kit/vela"
)

func define(r vela.Router) {
	r.GET("/api/v1/arr/agent/mem/status", xEnv.Then(func(ctx *fasthttp.RequestCtx) error {
		_G.Update()
		chunk, err := json.Marshal(_G)
		if err != nil {
			return err
		}
		ctx.Write(chunk)
		return nil
	}))
}
