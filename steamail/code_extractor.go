// This code is highly subject to steam's email format. It'll probably
// be ok, but may require some tinkering in the regexp in
// extractCode(). Not ideal, but simpler than parsing things.
package steamail

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/mail"
	"regexp"
)

type steamCodeExtractor struct{}

func NewSteamCodeExtractor() *steamCodeExtractor {
	return &steamCodeExtractor{}
}

func (e *steamCodeExtractor) ExtractCode(msg *mail.Message) (string, error) {
	log.Printf("Attempting to extract code from email")
	body, err := e.extractBody(msg)
	if err != nil {
		log.Printf("Couldn't read message body: %v", err)
		return "", fmt.Errorf("could not extract message body: %v", err)
	}
	return e.extractCode(body)
}

func (e *steamCodeExtractor) extractBody(msg *mail.Message) ([]byte, error) {
	return ioutil.ReadAll(msg.Body)
}

func (e *steamCodeExtractor) extractCode(body []byte) (string, error) {
	// The email has only one <h3>...</h3> block; the code is wrapped in
	// that.
	codeRegexp := regexp.MustCompile("<h3>.*</h3>")
	if !codeRegexp.Match(body) {
		log.Printf("No regexp match for code in email")
		return "", fmt.Errorf("no code in email")
	}
	codeBytes := codeRegexp.Find(body)
	// Eh, ugly, but whatever
	code := string(codeBytes[4 : len(codeBytes)-5])
	log.Printf("Extracted auth code '%s'", code)
	return code, nil
}
