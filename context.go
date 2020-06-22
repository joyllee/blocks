package blocks

import (
	"context"
	"net/http"
)

type HCtx struct {
	Ctx    context.Context    `json:"-"`
	Cancel context.CancelFunc `json:"-"`

	ResponseWriter http.ResponseWriter `json:"-"`
	Request        *http.Request       `json:"-"`
}

func NewHTTPContext() *HCtx {
	hCtx := &HCtx{}
	//hCtx.Ctx, hCtx.Cancel = context.WithCancel(signal.GetSignalContext().Ctx)
	//hCtx.SetTraceID(fmt.Sprintf("%s:0:0", utils.GetTraceID()))

	return hCtx
}
