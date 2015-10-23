package tweetmessage

import (
	"fmt"
	"message"
	"strings"
	"time"
)

var messageHistory []message.Message

type TweetMessage struct {
	// as long as you like
	// Print()
	//
	//Date: (<dateSent (2015/11/12 02:02:12pm))
	//From: <sender>
	//To: <recipient[0]>; <recipient[1]>; ...etc
	//Subject: <subject>
	//Body: <body>
	//Send() should add to a map slice called tweetHistory - if subject is the same then it should be organised like threads
	message.GenericMessage
}

func (e TweetMessage) Validate() (bool, error) {
	return true, nil
}

func (e *TweetMessage) Send() (bool, error) {
	e.DateSent = time.Now()
	messageHistory = append(messageHistory, message.Message(e))
	return true, nil
}

func (e *TweetMessage) Print() string {
	//@<sender> - <dateSent> (2015-11-12)
	//<body>
	//Mentions: @<mention[0]>, @<mention[1]>..etc
	r := ""
	r = r + fmt.Sprintf("@<%s> - <%s> (%s)\n", e.Sender, e.DateSent, e.DateSent)
	r = r + fmt.Sprintf("<%s>\n", e.Body)
	r = r + fmt.Sprintf("Mentions: ")

	// this sucks because i can't map @ in front of the strings so I have to do it twice... ergh
	mentions := make([]string, 0)
	for _, rcpt := range e.Recipients {
		mentions = append(mentions, fmt.Sprintf("@%s", rcpt))
	}
	r = r + fmt.Sprint(strings.Join(mentions, ","))
	r = r + "\n"
	return r

}

func New(sender, body string, recipients []string) (*TweetMessage, error) {
	em := TweetMessage{message.GenericMessage{Sender: sender, Body: body, Recipients: recipients}}
	return &em, nil
}

func (e *TweetMessage) GetHistory() ([]message.Message, error) {
	return messageHistory, nil
}
