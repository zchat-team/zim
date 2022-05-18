package typ

type MsgReadReceipt struct {
	Uin         string `json:"uin"`
	ReceiptTime int64  `json:"receipt_time"`
}

type MsgDeliveryReceipt struct {
	Uin         string `json:"uin"`
	ReceiptTime int64  `json:"receipt_time"`
}

type MsgRecall struct {
	Operator string `json:"operator"`
	Id       int64  `json:"id"`
}
