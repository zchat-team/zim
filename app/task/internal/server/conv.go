package server

import "fmt"

const (
	prefixConv        = "conv:%s:%s"          // conv:uin:peer
	prefixConvSync    = "conv_sync:%s"        // conv_sync:uin
	prefixConvMsgSync = "conv_msg_sync:%s:%s" // conv_msg_sync:uin:peer
	prefixConvMsg     = "conv_msg:%s:%s:%d"   // conv_msg:min:max:id 或 conv_msg:uin:peer:id
	prefixMsg         = "msg:%s:%d"           // msg:uin:id
	prefixMsgSync     = "msg_sync:%s"         // msg_sync:uin
)

func KeyConv(uin, peer string) string {
	return fmt.Sprintf(prefixConv, uin, peer)
}

func KeyConvSync(uin string) string {
	return fmt.Sprintf(prefixConvSync, uin)
}

func KeyConvMsgSync(uin, peer string) string {
	return fmt.Sprintf(prefixConvMsgSync, uin, peer)
}

func KeyConvMsg(first, second string, id int64) string {
	return fmt.Sprintf(prefixConvMsg, first, second, id)
}

func KeyMsg(uin string, id int64) string {
	return fmt.Sprintf(prefixMsg, uin, id)
}

func KeyMsgSync(uin string) string {
	return fmt.Sprintf(prefixMsgSync, uin)
}
