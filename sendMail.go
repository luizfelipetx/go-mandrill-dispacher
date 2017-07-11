package main

import (
	"fmt"
	"github.com/mostafah/mandrill"
)
func main() {
	mandrill.Key = "YOUR_KEY"
	msg := mandrill.NewMessageTo("email@to.com", "PersonName").AddRecipient("cc@to.com","PersonName")
	msg.HTML = ""  //if is html message put content here
	msg.Text = "" // optional
	msg.Subject = "Subject Text"
	msg.FromEmail = "from@from.com"
	msg.FromName = "PersonNameFrom"
	_, err := msg.Send(false)
	fmt.Println(err)
	fmt.Print("sending....")
}

