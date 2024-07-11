package menubotlib

import (
	"database/sql"
	"time"
)

type ConversationContext struct {
	UserInfo     UserInfo
	UserExisted  bool
	PriceList    []CatalogueSelection
	CurrentOrder CustomerOrder
	MessageBody  string
	DBReadTime   time.Time
}

func NewConversationContext(db *sql.DB, senderNumber, messagebody string, pricelist []CatalogueSelection, isAutoInc bool) *ConversationContext {
	userInfo, curOrder, userExisted := NewUserInfo(db, senderNumber, isAutoInc)
	context := &ConversationContext{
		UserInfo:     userInfo,
		UserExisted:  userExisted,
		PriceList:    pricelist,
		CurrentOrder: curOrder,
		MessageBody:  messagebody,
		DBReadTime:   time.Now(),
	}

	return context
}
