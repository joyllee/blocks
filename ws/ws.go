package ws

import (
	"github.com/joyllee/blocks"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

// 允许跨域
func SetCheckOrigin(f func(r *http.Request) bool) {
	upgrader.CheckOrigin = f
}

func NewUpgrader(f func(r *http.Request) bool) websocket.Upgrader {
	return websocket.Upgrader{
		CheckOrigin: f,
	}
}

type WsIns struct {
	ctx         *blocks.HCtx
	ws          *websocket.Conn
	keepTimeout time.Duration
	mt          *sync.Mutex
}

func NewWS(ctx *blocks.HCtx, h http.Header) (wsIns *WsIns, err error) {
	ws, err := upgrader.Upgrade(ctx.ResponseWriter, ctx.Request, h)
	if err != nil {
		return
	}
	wsIns = &WsIns{
		ctx:         ctx,
		keepTimeout: 60,
		mt:          new(sync.Mutex),
	}
	wsIns.ws = ws
	go wsIns.keep()

	return
}

func NewWSWithUpgrader(ctx *blocks.HCtx, h http.Header, upgrader websocket.Upgrader) (wsIns *WsIns, err error) {
	ws, err := upgrader.Upgrade(ctx.ResponseWriter, ctx.Request, h)
	if err != nil {
		return
	}
	wsIns = &WsIns{
		ctx:         ctx,
		keepTimeout: 60,
	}
	wsIns.ws = ws
	go wsIns.keep()

	return
}

func (wsIns *WsIns) Close() error {
	//wsIns.HTTPCtx.Cancel()
	return wsIns.ws.Close()
}

func (wsIns *WsIns) keep() {
FOR:
	for {
		select {
		case <-wsIns.ctx.Ctx.Done():
			//发送个信号给客户端，由客户端关闭
			wsIns.WriteCloseMessage(websocket.CloseServiceRestart, "keep ctx done")
			break FOR
		case <-time.After(wsIns.keepTimeout * time.Second):
			err := wsIns.WritePingMessage()
			if err != nil {
				wsIns.ctx.Warnf("keep error: %v", err)
				wsIns.ctx.Warn(wsIns.Close())
				break FOR
			}
		}
	}
}
func (wsIns *WsIns) ReadMessage() (messageType int, p []byte, err error) {
	return wsIns.ws.ReadMessage()
}

func (wsIns *WsIns) WritePingMessage() (err error) {
	wsIns.mt.Lock()
	defer wsIns.mt.Unlock()
	return wsIns.ws.WriteMessage(websocket.PingMessage, nil)
}

func (wsIns *WsIns) WriteTextMessage(data []byte) (err error) {
	wsIns.mt.Lock()
	defer wsIns.mt.Unlock()
	return wsIns.ws.WriteMessage(websocket.TextMessage, data)
}

func (wsIns *WsIns) WriteCloseMessage(closeCode int, text string) error {
	wsIns.mt.Lock()
	defer wsIns.mt.Unlock()
	return wsIns.ws.WriteMessage(websocket.CloseMessage,
		websocket.FormatCloseMessage(closeCode, text))
}

func (wsIns *WsIns) IsWebSocketCloseError(err error) bool {
	if err == nil {
		return false
	}
	close := []int{
		//正常关闭
		websocket.CloseNormalClosure,
		//当客户端页面刷新，ReadMessage就报这个错误
		websocket.CloseGoingAway,
		websocket.CloseAbnormalClosure,
		//发送CloseMessage给客户端后，服务器端就会收到这个
		websocket.CloseNoStatusReceived,
	}
	//服务器端或客户端close后，再对客户端写入，就会ErrCloseSent
	if websocket.IsCloseError(err, close...) || err == websocket.ErrCloseSent {
		wsIns.ctx.Infof("error: %v", err)
		return true
	}
	//服务器端close后，再对客户端读取，就会这个错误
	if strings.Contains(err.Error(), "closed network") {
		wsIns.ctx.Infof("error: %v", err)
		return true
	}

	wsIns.ctx.Warn(err)
	return false
}
