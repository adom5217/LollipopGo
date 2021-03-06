package MsgHandleClt

import (
	"golang.org/x/net/websocket"
)

type Msg_data interface {
	HandleCltProtocol(protocol interface{}, protocol2 interface{}, ProtocolData map[string]interface{}, Connection *websocket.Conn) interface{}
	HandleCltProtocol2(protocol2 interface{}, ProtocolData map[string]interface{}, Connection *websocket.Conn) interface{}
	PlayerSendMessage(senddata interface{}) int
}
