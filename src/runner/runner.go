package main

import (
	"emailmessage"
	"fmt"
	"log"
	"message"
	"smsmessage"
	"tweetmessage"
)

func main() {
	emailMessage, _ := emailmessage.New("sender person", "subject 1", "body thing 1 ", []string{"recpt1", "recpt2"})
	smsMessage, _ := smsmessage.New("sender person", "body thing 2", []string{"recpt1", "recpt2"})
	tweetMessage, _ := tweetmessage.New("sender person", "body thing 3", []string{"recpt1", "recpt2"})

	_, err := emailMessage.Send()
	_, err = smsMessage.Send()
	_, err = tweetMessage.Send()

	checkErr(err)

	collection := []message.Message{emailMessage, smsMessage, tweetMessage}
	for _, message := range collection {
		fmt.Println(message.Print())
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
