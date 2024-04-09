package view

import (
	"context"

	"github.com/Kotletta-TT/GophKeeper/proto"
	"github.com/rivo/tview"
)

type RegisterView struct {
	app   *tview.Application
	form  *tview.Form
	pages *tview.Pages
	api   proto.UserServiceClient
}

func NewRegisterView(client proto.UserServiceClient, app *tview.Application, pages *tview.Pages) *RegisterView {
	r := &RegisterView{
		api:   client,
		app:   app,
		pages: pages,
	}
	r.NewRegisterForm()
	return r
}

func (r *RegisterView) GetForm() *tview.Form {
	return r.form
}

func (r *RegisterView) NewRegisterForm() {
	r.form = tview.NewForm()
	r.form.AddInputField("Username", "", 20, nil, nil).
		AddPasswordField("Password", "", 10, '*', nil).
		AddButton("Register", func() {
			var modal *tview.Modal
			usernameField := r.form.GetFormItemByLabel("Username").(*tview.InputField)
			passwordField := r.form.GetFormItemByLabel("Password").(*tview.InputField)
			username := usernameField.GetText()
			password := passwordField.GetText()
			resp, err := r.api.CreateUser(context.TODO(), &proto.UserRequest{Login: username, Password: password})
			if err != nil {
				modal = tview.NewModal().SetText(err.Error()).AddButtons([]string{"OK"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
					r.pages.RemovePage("Modal")
					r.pages.SwitchToPage("Register")
				})
				r.app.SetRoot(r.pages, true)
				r.pages.AddPage("Modal", modal, true, true)
				r.pages.SwitchToPage("Modal")
			} else if resp.Error != "" {
				modal = tview.NewModal().SetText(resp.Error).AddButtons([]string{"OK"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
					r.pages.RemovePage("Modal")
					r.pages.SwitchToPage("Register")
				})
				r.app.SetRoot(r.pages, true)
				r.pages.AddPage("Modal", modal, true, true)
				r.pages.SwitchToPage("Modal")
			} else {
				r.pages.SwitchToPage("Register")
			}

			// Здесь вы можете обработать ответ от сервера
		}).
		SetButtonsAlign(tview.AlignCenter)
}
