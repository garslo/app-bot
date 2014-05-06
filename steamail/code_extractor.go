// This code is highly subject to steam's email format. It'll probably
// be ok, but may require some tinkering in the regexp in
// extractCode(). Not ideal, but simpler than parsing things.
package steamail

import (
	"fmt"
	"io/ioutil"
	"net/mail"
	"regexp"
)

type steamCodeExtractor struct{}

func NewSteamCodeExtractor() *steamCodeExtractor {
	return &steamCodeExtractor{}
}

func (e *steamCodeExtractor) ExtractCode(msg *mail.Message) (string, error) {
	body, err := e.extractBody(msg)
	if err != nil {
		return "", fmt.Errorf("could not extract message body: %v", err)
	}
	return e.extractCode(body), nil
}

func (e *steamCodeExtractor) extractBody(msg *mail.Message) ([]byte, error) {
	return ioutil.ReadAll(msg.Body)
}

func (e *steamCodeExtractor) extractCode(body []byte) string {
	// The email has only one <h3>...</h3> block; the code is wrapped in
	// that.
	codeRegexp := regexp.MustCompile("<h3>.*</h3>")
	codeBytes := codeRegexp.Find(body)
	// Eh, ugly, but whatever
	return string(codeBytes[4 : len(codeBytes)-5])
}
