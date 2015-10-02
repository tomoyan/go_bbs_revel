package models

import (
	//"fmt"
	//"github.com/go-gorp/gorp"
	"github.com/revel/revel"
	//"time"
)

type Message struct {
	MessageId int
	Name      string
	Email     string
	Title     string
	Message   string
	Created   string
}

// Interface for Validate() and
// validation can pass in the key prefix ("message.")
func (message Message) Validate(v *revel.Validation) {

	v.Check(message.Name,
		revel.Required{},
		revel.MinSize{3},
	).Message("Name is required and more than 3 characters!")
	//v.Required(message.Name).Message("Name is required!")
	//v.MinSize(message.Name, 3).Message("Name is not long enough!")
	v.Required(message.Email).Message("Email is required!")
	v.Required(message.Title).Message("Title is required!")
	v.Required(message.Message).Message("Message is required!")
}
