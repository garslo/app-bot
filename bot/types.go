package bot

import "github.com/Philipp15b/go-steam"

type BotConfig struct {
	SentryFile string
	Username   string
	Password   string
}

type Bot interface {
	Login()
}

type EventProxy interface {
	StartProxying()
	GetChatMsgEventChan() chan *steam.ChatMsgEvent
	GetConnectedEventChan() chan *steam.ConnectedEvent
	GetFriendListEventChan() chan *steam.FriendListEvent
	GetFriendStateEventChan() chan *steam.FriendStateEvent
	GetGroupStateEventChan() chan *steam.GroupStateEvent
	GetLoggedOnEventChan() chan *steam.LoggedOnEvent
	GetMachineAuthUpdateEventChan() chan *steam.MachineAuthUpdateEvent
	GetTradeProposedEventChan() chan *steam.TradeProposedEvent
	GetTradeResultEventChan() chan *steam.TradeResultEvent
	GetTradeSessionStartEventChan() chan *steam.TradeSessionStartEvent
	GetWebLoggedOnEventChan() chan *steam.WebLoggedOnEvent
	GetWebSessionIdEventChan() chan *steam.WebSessionIdEvent
}
