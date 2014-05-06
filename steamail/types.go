package steamail

import "net/mail"

type SteamCodeExtractor interface {
	ExtractCode(msg *mail.Message) (string, error)
}

type SteamCodeGetter interface {
	GetSteamCode() (string, error)
}
