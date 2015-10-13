package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"github.com/tomoyan/go_bbs_revel/app/models"
	"time"
)

type BBS struct {
	*revel.Controller
	GorpController
}

func (c BBS) Index() revel.Result {
	revel.INFO.Println("Start BBS.Index")
	results, err := c.Txn.Select(models.Message{},
		`select * from Message order by MessageId desc`)
	if err != nil {
		panic(err)
	}

	// Filling Message Slice
	var messages []*models.Message
	for _, r := range results {
		b := r.(*models.Message)
		messages = append(messages, b)
	}
	return c.Render(messages)
}

func (c BBS) ConfirmCreate(message models.Message) revel.Result {
	revel.INFO.Println("Start BBS.ConfirmCreate")
	message.Created = time.Now().Format(DATE_FORMAT)

	fmt.Println("Input Message:", message)
	//fmt.Println("### Form Input ###")
	//	fmt.Println("Name:", message.Name)
	//	fmt.Println("Email:", message.Email)
	//	fmt.Println("Title:", message.Title)
	//	fmt.Println("Message:", message.Message)
	//	fmt.Println("Created:", message.Created)

	//	c.Validation.Required(message.Name).Message("Name is required!")
	//	c.Validation.MinSize(message.Name, 3).Message("Name is not long enough!")
	//	c.Validation.Required(message.Email).Message("Email is required!")
	//	c.Validation.Required(message.Title).Message("Title is required!")
	//	c.Validation.Required(message.Message).Message("Message is required!")

	message.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(BBS.Index)
	}

	err := c.Txn.Insert(&message)
	if err != nil {
		panic(err)
	}

	//c.Flash.Success("Thank you, %sさん!", message.Name)
	return c.Render()
}
