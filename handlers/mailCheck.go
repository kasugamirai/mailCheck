package handlers

import (
	"strings"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

func ConnectToEmail(email, password string) (*client.Client, error) {
	c, err := client.DialTLS("imap.gmail.com:993", nil)
	if err != nil {
		return nil, err
	}
	if err := c.Login(email, password); err != nil {
		return nil, err
	}
	return c, nil
}

func CheckForContents(c *client.Client) ([]string, error) {
	if _, err := c.Select("INBOX", false); err != nil {
		return nil, err
	}

	criteria := imap.NewSearchCriteria()
	criteria.WithoutFlags = []string{"\\Seen"}
	uids, err := c.Search(criteria)
	if err != nil {
		return nil, err
	}

	seqset := new(imap.SeqSet)
	seqset.AddNum(uids...)

	messages := make(chan *imap.Message, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope}, messages)
	}()

	var foundEmails []string
	for msg := range messages {
		subject := msg.Envelope.Subject
		if strings.Contains(subject, "123") {
			foundEmails = append(foundEmails, subject)
		}
	}

	if err := <-done; err != nil {
		return nil, err
	}

	return foundEmails, nil
}
