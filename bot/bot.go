package bot

import (
	"crypto/sha1"
	"io/ioutil"
	"log"

	"github.com/Philipp15b/go-steam"
	"github.com/garslo/app-bot/steamail"
)

type bot struct {
	client     *steam.Client
	codeGetter steamail.SteamCodeGetter
	config     *BotConfig
	eventProxy EventProxy
}

func NewBot(client *steam.Client, codeGetter steamail.SteamCodeGetter, config *BotConfig) *bot {
	b := &bot{
		client:     client,
		codeGetter: codeGetter,
		config:     config,
		eventProxy: NewEventProxy(client.Events()),
	}
	go b.handleEvents()
	return b
}

func (b *bot) handleEvents() {
	go b.eventProxy.StartProxying()
	go b.handleFriendRequests()
}

func (b *bot) handleFriendRequests() {
	for event := range b.eventProxy.GetFriendStateEventChan() {
		log.Printf("Got a FriendStateEvent:")
		log.Printf("IsFriend(): %v", event.IsFriend())
		log.Printf("IsMember(): %v", event.IsMember())
		log.Printf("SteamId: %v", event.SteamId)
		log.Printf("Relationship: %v", event.Relationship)
		log.Printf("Adding as friend")
		b.client.Social.AddFriend(event.SteamId)
	}
}

func (b *bot) Login() {
	logOnDetails := b.getLogOnDetails()
	b.login(logOnDetails)
}

func (b *bot) login(details *steam.LogOnDetails) {
	b.client.Auth.LogOn(details)
}

func (b *bot) getLogOnDetails() *steam.LogOnDetails {
	// These fellas might have a value, or they might be empty. Steam
	// will give us a response later.
	authCode, sentryHash := b.getCodeOrHash()
	return &steam.LogOnDetails{
		Username:       b.config.Username,
		Password:       b.config.Password,
		AuthCode:       authCode,
		SentryFileHash: sentryHash,
	}
}

func (b *bot) getCodeOrHash() (string, []byte) {
	sentryHash, err := b.getSentryHash()
	if err == nil {
		return "", sentryHash
	}
	authCode, err := b.codeGetter.GetSteamCode()
	if err == nil {
		return authCode, []byte{}
	}
	return "", []byte{}
}

func (b *bot) getSentryHash() ([]byte, error) {
	sentryData, err := ioutil.ReadFile(b.config.SentryFile)
	if err != nil {
		return []byte{}, err
	}
	h := sha1.New()
	h.Write(sentryData)
	return h.Sum(nil), nil
}
