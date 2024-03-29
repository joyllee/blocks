package ws

import (
	"encoding/json"
	"github.com/joyllee/blocks/utils"
	"net/http"
	"testing"
	"time"
)

func TestWsInit(t *testing.T) {
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe("0.0.0.0:7777", nil)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := blocks.NewHTTPContext()
	ctx.ResponseWriter = w
	ctx.Request = r

	//允许跨域
	SetCheckOrigin(func(r *http.Request) bool {
		return true
	})
	wsIns, err := NewWS(ctx, nil)
	if err != nil {
		return
	}
	defer wsIns.Close()

	go func() {
		defer func() {
			if err := recover(); err != nil {
				wsIns.ctx.Warn(err, string(utils.GetStack()))
			}
		}()
		for {
			select {
			case <-ctx.Ctx.Done():
				return
			default:
				_, _, err := wsIns.ReadMessage()
				if wsIns.IsWebSocketCloseError(err) {
					ctx.Cancel()
					return
				}
			}
		}
	}()

	marshal, _ := json.Marshal("test")
	for {
		select {
		case <-ctx.Ctx.Done():
			ctx.Info("ws is done")
			return
		default:
			wsIns.WriteTextMessage(marshal)
			time.Sleep(5 * time.Second)
		}
	}

}
