package view

import (
	"context"

	"github.com/Kotletta-TT/GophKeeper/proto"
	"github.com/rivo/tview"
)

type LoginView struct {
	app   *tview.Application
	pages *tview.Pages
	api   proto.UserServiceClient
	form  *tview.Form
	token *string
}

func NewLoginView(client proto.UserServiceClient, app *tview.Application, pages *tview.Pages, token *string) *LoginView {
	l := &LoginView{
		api:   client,
		app:   app,
		pages: pages,
		token: token,
	}
	l.NewLoginForm()
	return l
}

func (l *LoginView) GetForm() *tview.Form {
	return l.form
}

func (l *LoginView) NewLoginForm() {
	l.form = tview.NewForm()
	l.form.AddInputField("Username", "", 20, nil, nil).
		AddPasswordField("Password", "", 10, '*', nil).
		AddButton("Login", func() {
			var modal *tview.Modal
			usernameField := l.form.GetFormItemByLabel("Username").(*tview.InputField)
			passwordField := l.form.GetFormItemByLabel("Password").(*tview.InputField)
			username := usernameField.GetText()
			password := passwordField.GetText()
			resp, err := l.api.Login(context.TODO(), &proto.UserRequest{Login: username, Password: password})
			if err != nil {
				modal = tview.NewModal().SetText(err.Error()).AddButtons([]string{"OK"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
					l.pages.RemovePage("Modal")
					l.pages.SwitchToPage("Login")
				})
				l.app.SetRoot(l.pages, true)
				l.pages.AddPage("Modal", modal, true, true)
				l.pages.SwitchToPage("Modal")
			} else {
				*l.token = resp.Token
				modal = tview.NewModal().SetText(resp.Token).AddButtons([]string{"OK"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
					l.pages.RemovePage("Modal")
					l.pages.SwitchToPage("Card")
				})
				// l.app.SetRoot(l.pages, true)
				l.pages.AddPage("Modal", modal, true, true)
				l.pages.SwitchToPage("Modal")
			}
		}).
		AddButton("Exit", func() {
			l.app.Stop()
		})
}
