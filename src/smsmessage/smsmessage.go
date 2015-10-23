package smsmessage

import (
	"fmt"
	"message"
	"time"
)

var messageHistory []message.Message

type SMSMessage struct {
	// as long as you like
	// Print()
	//
	//Date: (<dateSent (2015/11/12 02:02:12pm))
	//From: <sender>
	//To: <recipient[0]>; <recipient[1]>; ...etc
	//Subject: <subject>
	//Body: <body>
	//Send() should add to a map slice called smsHistory - if subject is the same then it should be organised like threads
	message.GenericMessage
}

func (e SMSMessage) Validate() (bool, error) {
	return true, nil
}

func (e *SMSMessage) Send() (bool, error) {
	e.DateSent = time.Now()
	messageHistory = append(messageHistory, message.Message(e))
	return true, nil
}

func (e *SMSMessage) Print() string {
	//<sender> -> <body> (<dateSent (2015/11/12 14:02:12pm))
	return fmt.Sprintf("<%s> -> <%s> (<%v>)", e.Sender, e.Body, e.DateSent)

}

func New(sender, body string, recipients []string) (*SMSMessage, error) {
	em := SMSMessage{message.GenericMessage{Sender: sender, Body: body, Recipients: recipients}}
	return &em, nil
}

func (e *SMSMessage) GetHistory() ([]message.Message, error) {
	return messageHistory, nil
}
