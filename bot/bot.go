package bot

import (
	"crypto/sha1"
	"io/ioutil"
	"log"
	"time"

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
	go b.handleErrors()
	go b.handleChat()
}

func (b *bot) handleChat() {
	for chatEvent := range b.eventProxy.GetChatMsgEventChan() {
		log.Printf("Got chat message: %s", chatEvent.Message)
		log.Printf("Chatroom: %v", chatEvent.Chatroom)
		log.Printf("Sender: %v", chatEvent.Sender)
	}
}

func (b *bot) handleFriendRequests() {
	for event := range b.eventProxy.GetFriendStateEventChan() {
		log.Printf("Got a FriendStateEvent:")
		log.Printf("IsFriend(): %v", event.IsFriend())
		log.Printf("IsMember(): %v", event.IsMember())
		log.Printf("SteamId: %v", event.SteamId)
		log.Printf("Relationship: %v", event.Relationship)
		log.Printf("Adding as friend")
		if !event.IsFriend() {
			b.client.Social.AddFriend(event.SteamId)
		}
	}
}

func (b *bot) handleErrors() {
	for {
		select {
		case fatalErr := <-b.eventProxy.GetFatalErrorChan():
			log.Printf("Got fatal error, going to reconnect: %v", fatalErr)
			// TODO: reconnection shiz
			go b.Login()
		case err := <-b.eventProxy.GetErrorChan():
			log.Printf("Got an error: %v", err)
			// TODO: handling shiz
		}
	}
}

func (b *bot) Login() {
	logOnDetails := b.getLogOnDetails()
	b.login(logOnDetails)
}

func (b *bot) login(details *steam.LogOnDetails) {
	b.client.Auth.LogOn(details)
	for {
		select {
		case <-b.eventProxy.GetLoggedOnEventChan():
			log.Printf("Login successful")
			return
		case <-time.After(b.config.LoginRetryInterval):
			log.Printf("Retrying login...")
			b.client.Auth.LogOn(details)
		}
	}
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
	log.Printf("Fetching sentry hash from '%s'", b.config.SentryFile)
	sentryData, err := ioutil.ReadFile(b.config.SentryFile)
	if err != nil {
		log.Printf("Failed to get sentry hash")
		return []byte{}, err
	}
	log.Printf("Got sentry hash (unprintable)")
	h := sha1.New()
	h.Write(sentryData)
	return h.Sum(nil), nil
}
