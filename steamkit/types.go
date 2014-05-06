package steamkit

import (
	"github.com/Philipp15b/go-steam"
	"github.com/Philipp15b/go-steam/internal"
	"github.com/Philipp15b/go-steam/steamid"
)

type Auth interface {
	LogOn(*steam.LogOnDetails)
	HandlePacket(*internal.PacketMsg)
}

func NewAuth(auth *steam.Auth) Auth {
	return auth
}

type Social interface {
	HandlePacket(*internal.PacketMsg)
	SetPersonaState(internal.EPersonaState)
	SendChatMessage(steamid.SteamId, string)
	SendMessage(steamid.SteamId, string, internal.EChatEntryType)
	AddFriend(steamid.SteamId)
	RemoveFriend(steamid.SteamId)
	GetFriendsList() *steam.FriendsList
	GetGroupsList() *steam.GroupsList
}

type social struct {
	*steam.Social
}

func NewSocial(s *steam.Social) Social {
	return &social{
		Social: s,
	}
}

func (s *social) GetFriendsList() *steam.FriendsList {
	return s.Friends
}

func (s *social) GetGroupsList() *steam.GroupsList {
	return s.Groups
}

type Trader interface {
	HandlePacket(*internal.PacketMsg)
	RequestTrade(steamid.SteamId)
	RespondRequest(steam.TradeRequestId, bool)
	CancelRequest(steamid.SteamId)
}

func NewTrader(trading *steam.Trading) Trader {
	return trading
}

type Weber interface {
	HandlePacket(*internal.PacketMsg)
	LogOn()
	GetWebSessionId() string
	GetSteamLogin() string
}

type weber struct {
	*steam.Web
}

func NewWeber(w *steam.Web) Weber {
	return &weber{
		Web: w,
	}
}

func (w *weber) GetWebSessionId() string {
	return w.WebSessionId
}

func (w *weber) GetSteamLogin() string {
	return w.SteamLogin
}
