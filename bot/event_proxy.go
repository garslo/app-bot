package bot

import (
	"log"

	"github.com/Philipp15b/go-steam"
)

type eventProxy struct {
	eventChan                  <-chan interface{}
	ChatMsgEventChan           chan *steam.ChatMsgEvent
	ConnectedEventChan         chan *steam.ConnectedEvent
	FriendListEventChan        chan *steam.FriendListEvent
	FriendStateEventChan       chan *steam.FriendStateEvent
	GroupStateEventChan        chan *steam.GroupStateEvent
	LoggedOnEventChan          chan *steam.LoggedOnEvent
	MachineAuthUpdateEventChan chan *steam.MachineAuthUpdateEvent
	TradeProposedEventChan     chan *steam.TradeProposedEvent
	TradeResultEventChan       chan *steam.TradeResultEvent
	TradeSessionStartEventChan chan *steam.TradeSessionStartEvent
	WebLoggedOnEventChan       chan *steam.WebLoggedOnEvent
	WebSessionIdEventChan      chan *steam.WebSessionIdEvent
	FatalErrorChan             chan *steam.FatalError
	ErrorChan                  chan error
}

func NewEventProxy(eventChan <-chan interface{}) EventProxy {
	return &eventProxy{
		eventChan: eventChan,
	}
}

func (p *eventProxy) StartProxying() {
	for event := range p.eventChan {
		switch e := event.(type) {
		case *steam.ChatMsgEvent:
			if p.ChatMsgEventChan == nil {
				log.Print("Ignoring ChatMsgEvent")
				break
			}
			p.ChatMsgEventChan <- e
		case *steam.ConnectedEvent:
			if p.ConnectedEventChan == nil {
				log.Print("Ignoring ConnectedEvent")
				break
			}
			p.ConnectedEventChan <- e
		case *steam.FriendListEvent:
			if p.FriendListEventChan == nil {
				log.Print("Ignoring FriendListEvent")
				break
			}
			p.FriendListEventChan <- e
		case *steam.FriendStateEvent:
			if p.FriendStateEventChan == nil {
				log.Print("Ignoring FriendStateEvent")
				break
			}
			p.FriendStateEventChan <- e
		case *steam.GroupStateEvent:
			if p.GroupStateEventChan == nil {
				log.Print("Ignoring GroupStateEvent")
				break
			}
			p.GroupStateEventChan <- e
		case *steam.LoggedOnEvent:
			if p.LoggedOnEventChan == nil {
				log.Print("Ignoring LoggedOnEvent")
				break
			}
			p.LoggedOnEventChan <- e
		case *steam.MachineAuthUpdateEvent:
			if p.MachineAuthUpdateEventChan == nil {
				log.Print("Ignoring MachineAuthUpdateEvent")
				break
			}
			p.MachineAuthUpdateEventChan <- e
		case *steam.TradeProposedEvent:
			if p.TradeProposedEventChan == nil {
				log.Print("Ignoring TradeProposedEvent")
				break
			}
			p.TradeProposedEventChan <- e
		case *steam.TradeResultEvent:
			if p.TradeResultEventChan == nil {
				log.Print("Ignoring TradeResultEvent")
				break
			}
			p.TradeResultEventChan <- e
		case *steam.TradeSessionStartEvent:
			if p.TradeSessionStartEventChan == nil {
				log.Print("Ignoring TradeSessionStartedEvent")
				break
			}
			p.TradeSessionStartEventChan <- e
		case *steam.WebLoggedOnEvent:
			if p.WebLoggedOnEventChan == nil {
				log.Print("Ignoring WebLoggedOnEvent")
				break
			}
			p.WebLoggedOnEventChan <- e
		case *steam.WebSessionIdEvent:
			if p.WebSessionIdEventChan == nil {
				log.Print("Ignoring WebSessionIdEvent")
				break
			}
			p.WebSessionIdEventChan <- e
		default:
			log.Printf("Unknown event: %v", e)
		}
	}
}

func (p *eventProxy) GetChatMsgEventChan() chan *steam.ChatMsgEvent {
	if p.ChatMsgEventChan == nil {
		p.ChatMsgEventChan = make(chan *steam.ChatMsgEvent)
	}
	return p.ChatMsgEventChan
}

func (p *eventProxy) GetConnectedEventChan() chan *steam.ConnectedEvent {
	if p.ConnectedEventChan == nil {
		p.ConnectedEventChan = make(chan *steam.ConnectedEvent)
	}
	return p.ConnectedEventChan
}

func (p *eventProxy) GetFriendListEventChan() chan *steam.FriendListEvent {
	if p.FriendListEventChan == nil {
		p.FriendListEventChan = make(chan *steam.FriendListEvent)
	}
	return p.FriendListEventChan
}

func (p *eventProxy) GetFriendStateEventChan() chan *steam.FriendStateEvent {
	if p.FriendStateEventChan == nil {
		p.FriendStateEventChan = make(chan *steam.FriendStateEvent)
	}
	return p.FriendStateEventChan
}

func (p *eventProxy) GetGroupStateEventChan() chan *steam.GroupStateEvent {
	if p.GroupStateEventChan == nil {
		p.GroupStateEventChan = make(chan *steam.GroupStateEvent)
	}
	return p.GroupStateEventChan
}

func (p *eventProxy) GetLoggedOnEventChan() chan *steam.LoggedOnEvent {
	if p.LoggedOnEventChan == nil {
		p.LoggedOnEventChan = make(chan *steam.LoggedOnEvent)
	}
	return p.LoggedOnEventChan
}

func (p *eventProxy) GetMachineAuthUpdateEventChan() chan *steam.MachineAuthUpdateEvent {
	if p.MachineAuthUpdateEventChan == nil {
		p.MachineAuthUpdateEventChan = make(chan *steam.MachineAuthUpdateEvent)
	}
	return p.MachineAuthUpdateEventChan
}

func (p *eventProxy) GetTradeProposedEventChan() chan *steam.TradeProposedEvent {
	if p.TradeProposedEventChan == nil {
		p.TradeProposedEventChan = make(chan *steam.TradeProposedEvent)
	}
	return p.TradeProposedEventChan
}

func (p *eventProxy) GetTradeResultEventChan() chan *steam.TradeResultEvent {
	if p.TradeResultEventChan == nil {
		p.TradeResultEventChan = make(chan *steam.TradeResultEvent)
	}
	return p.TradeResultEventChan
}

func (p *eventProxy) GetTradeSessionStartEventChan() chan *steam.TradeSessionStartEvent {
	if p.TradeSessionStartEventChan == nil {
		p.TradeSessionStartEventChan = make(chan *steam.TradeSessionStartEvent)
	}
	return p.TradeSessionStartEventChan
}

func (p *eventProxy) GetWebLoggedOnEventChan() chan *steam.WebLoggedOnEvent {
	if p.WebLoggedOnEventChan == nil {
		p.WebLoggedOnEventChan = make(chan *steam.WebLoggedOnEvent)
	}
	return p.WebLoggedOnEventChan
}

func (p *eventProxy) GetWebSessionIdEventChan() chan *steam.WebSessionIdEvent {
	if p.WebSessionIdEventChan == nil {
		p.WebSessionIdEventChan = make(chan *steam.WebSessionIdEvent)
	}
	return p.WebSessionIdEventChan
}

func (p *eventProxy) GetFatalErrorChan() chan *steam.FatalError {
	if p.FatalErrorChan == nil {
		p.FatalErrorChan = make(chan *steam.FatalError)
	}
	return p.FatalErrorChan
}

func (p *eventProxy) GetErrorChan() chan error {
	if p.ErrorChan == nil {
		p.ErrorChan = make(chan error)
	}
	return p.ErrorChan
}
