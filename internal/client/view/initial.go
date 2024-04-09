package view

import (
	"github.com/Kotletta-TT/GophKeeper/proto"
	"github.com/rivo/tview"
)

type InitialView struct {
	app   *tview.Application
	pages *tview.Pages
	api   proto.UserServiceClient
	list  *tview.List
}

func NewInitialView(client proto.UserServiceClient, app *tview.Application, pages *tview.Pages) *InitialView {
	i := &InitialView{
		app:   app,
		pages: pages,
		api:   client,
	}
	i.NewList()
	return i
}

func (in *InitialView) NewList() {
	in.list = tview.NewList().
		AddItem("Login", "", 'l', nil).
		AddItem("Register", "", 'r', nil)
		// AddItem("GetCards", "", 'g', nil)

	in.list.SetSelectedFunc(func(i int, name string, s string, t rune) {
		in.pages.SwitchToPage(name)
	})
}

func (i *InitialView) GetList() *tview.List {
	return i.list
}
