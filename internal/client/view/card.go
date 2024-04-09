package view

import (
	"github.com/Kotletta-TT/GophKeeper/proto"
	"github.com/rivo/tview"
)

type CardView struct {
	list  *tview.List
	app   *tview.Application
	pages *tview.Pages
	api   proto.SecretCardServiceClient
	token *string
	usrID *string
	ccv   *CreateCardView
}

func NewCardView(card proto.SecretCardServiceClient, app *tview.Application, pages *tview.Pages, token, uID *string, fApi proto.FileSecretCardServiceClient) *CardView {
	cv := &CardView{
		api:   card,
		app:   app,
		pages: pages,
		token: token,
		usrID: uID,
		ccv:   NewCreateCardView(app, pages, card, fApi, token, uID),
	}
	cv.InitializeList()
	return cv
}

func (c *CardView) GetList() *tview.List {
	return c.list
}

func (c *CardView) InitializeList() {
	c.list = tview.NewList()
	c.list.AddItem("Create Card", "", 'c', nil)
	c.list.AddItem("Read Card", "", 'r', nil)
	c.list.AddItem("Update Card", "", 'u', nil)
	c.list.AddItem("Delete Card", "", 'd', nil)
	c.list.AddItem("List Cards", "", 'l', nil)
	c.list.AddItem("Sync", "", 's', nil)
	c.list.SetSelectedFunc(func(i int, name string, s string, t rune) {
		switch name {
		case "Create Card":
			c.pages.AddPage("CreateCard", c.ccv.GetForm(), true, true)
			c.pages.SwitchToPage("CreateCard")
		}
	})
}
