package message

import "time"

//type Message interface {
////	New() (m *Message, err error)
//	Send() (ok bool, err error)
//	Print() string
//	GetHistory() (history []interface{}, err error)
//}
type Message interface {
	Validate() (ok bool, err error)
	Send() (ok bool, err error)
	Print() string
	GetHistory() (history []Message, err error)
}

type GenericMessage struct {
	Sender     string
	Body       string
	DateSent   time.Time
	Recipients []string
}

//type EmailMessage struct {
// as long as you like
// Print()
//
//Date: (<dateSent (2015/11/12 02:02:12pm))
//From: <sender>
//To: <recipient[0]>; <recipient[1]>; ...etc
//Subject: <subject>
//Body: <body>
//Send() should add to a map slice called emailHistory - if subject is the same then it should be organised like threads
//	GenericMessage
//}

type SMSMessage struct {
	// 300 chars max
	// Print() should print like...
	//
	//<sender> -> <body> (<dateSent (2015/11/12 14:02:12pm))
	//
	//Send() should add to map slice called smsHistory organised by recievers
	GenericMessage
}

type TweetMessage struct {
	// 140 chars max
	//Print() should print like...
	//
	//@<sender> - <dateSent> (2015-11-12)
	//<body>
	//Mentions: @<mention[0]>, @<mention[1]>..etc

	//Send() should add to map slice called tweetHistory in date order
	// # mentions are optional
	GenericMessage
}

// Message funcs
//func (m Message) New() (*GenericMessage, error) {
//}

//func (m Message) Send() (bool, error) {
//}

//func (m Message) Print() string {
//}

//func (m Message) GetHistory() ([]interface{}, error) {
//}

//func (EmailMessage) New(sender, body string, recipients []string) (*GenericMessage, error) {
//	var gm = GenericMessage{
//		sender:     sender,
//		body:       body,
//		recipients: recipients,
//	}

//m.GenericMessage = gm

//	return &gm, nil
//}
