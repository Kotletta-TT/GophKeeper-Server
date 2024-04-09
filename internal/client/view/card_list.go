package view

import (
	"context"

	"github.com/Kotletta-TT/GophKeeper/proto"
	"github.com/rivo/tview"
)

type CardListView struct {
	app         *tview.Application
	pages       *tview.Pages
	api         proto.SecretCardServiceClient
	list        *tview.List
	initialized bool
}

func NewCardListView(card proto.SecretCardServiceClient, app *tview.Application, pages *tview.Pages) *CardListView {
	c := &CardListView{
		app:         app,
		pages:       pages,
		api:         card,
		initialized: false,
	}
	return c
}

func (c *CardListView) GetList() *tview.List {
	// Проверяем, была ли уже произведена инициализация списка
	if !c.initialized {
		// Создаем страницу с пустым списком
		c.list = tview.NewList()
		c.list.AddItem("Loading...", "", 0, nil)
		c.initialized = true
	} else {
		c.InitializeList()
	}
	return c.list
}

func (c *CardListView) InitializeList() {
	c.list.Clear()
	resp, err := c.api.ListSecretCard(context.Background(), &proto.ListSecretCardRequest{
		UserId: "6952cfcc-a2c2-46ca-870b-643f8e213ad6",
	})
	if err != nil {
		modal := tview.NewModal().SetText(err.Error()).AddButtons([]string{"OK"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			c.pages.RemovePage("Modal")
			c.pages.SwitchToPage("CardList")
		})
		c.app.SetRoot(c.pages, true)
		c.pages.AddPage("Modal", modal, true, true)
		c.pages.SwitchToPage("Modal")
		return
	}
	c.list = tview.NewList()
	cards := resp.GetCards()
	if cards == nil {
		c.list.AddItem("No cards", "", 'n', nil)
		c.list.SetSelectedFunc(func(i int, name string, s string, t rune) {
			c.pages.SwitchToPage("Initial")
		})
		return
	}
	for _, v := range cards {
		c.list.AddItem(v.Name, "", rune(v.Name[0]), nil)
	}
	c.list.SetSelectedFunc(func(i int, name string, s string, t rune) {
		c.pages.SwitchToPage(name)
	})
}
