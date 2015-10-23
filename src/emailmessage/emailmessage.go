package emailmessage

import (
	"fmt"
	"message"
	"time"
)

var messageHistory []message.Message

type EmailMessage struct {
	// as long as you like
	// Print()
	//
	//Date: (<dateSent (2015/11/12 02:02:12pm))
	//From: <sender>
	//To: <recipient[0]>; <recipient[1]>; ...etc
	//Subject: <subject>
	//Body: <body>
	//Send() should add to a map slice called emailHistory - if subject is the same then it should be organised like threads
	message.GenericMessage
	Subject string
}

func (e EmailMessage) Validate() (bool, error) {
	return true, nil
}

func (e *EmailMessage) Send() (bool, error) {
	e.DateSent = time.Now()
	messageHistory = append(messageHistory, message.Message(e))
	return true, nil
}

func (e *EmailMessage) Print() string {
	//Date: (<dateSent (2015/11/12 02:02:12pm))
	//From: <sender>
	//To: <recipient[0]>; <recipient[1]>; ...etc
	//Subject: <subject>
	//Body: <body>
	r := ""
	r = r + fmt.Sprintf("Date: (<%s>)\n", e.DateSent)
	r = r + fmt.Sprintf("From: <%s>\n", e.Sender)
	r = r + "To "
	for _, rcpt := range e.Recipients {
		r = r + fmt.Sprintf("%s;", rcpt)
	}
	r = r + "\n"
	r = r + fmt.Sprintf("Subject: %s\n", e.Subject)
	r = r + fmt.Sprintf("Body: %s\n", e.Body)
	return r

}

func New(sender, subject, body string, recipients []string) (*EmailMessage, error) {
	em := EmailMessage{message.GenericMessage{Sender: sender, Body: body, Recipients: recipients}, subject}
	return &em, nil
}

func (e *EmailMessage) GetHistory() ([]message.Message, error) {
	return messageHistory, nil
}
