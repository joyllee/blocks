package blocks

import (
	"context"
	"github.com/joyllee/blocks/logger"
	"github.com/sirupsen/logrus"
	"net/http"
)

type HCtx struct {
	Ctx    context.Context    `json:"-"`
	Cancel context.CancelFunc `json:"-"`

	ResponseWriter http.ResponseWriter `json:"-"`
	Request        *http.Request       `json:"-"`

	*logrus.Logger
}

func NewHTTPContext() *HCtx {
	hCtx := &HCtx{}
	hCtx.Ctx, hCtx.Cancel = context.WithCancel(context.Background())
	hCtx.Logger = logger.InitLogger()

	return hCtx
}
