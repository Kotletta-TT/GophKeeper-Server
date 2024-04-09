package view

import (
	"context"

	"github.com/Kotletta-TT/GophKeeper/proto"
	"github.com/rivo/tview"
	"google.golang.org/grpc/metadata"
)

type CreateCardView struct {
	form  *tview.Form
	api   proto.SecretCardServiceClient
	pages *tview.Pages
	app   *tview.Application
	token *string
	usrID *string
	fForm *CreateFileView
}

func NewCreateCardView(app *tview.Application, pages *tview.Pages, api proto.SecretCardServiceClient, fApi proto.FileSecretCardServiceClient, token, uID *string) *CreateCardView {
	return &CreateCardView{
		api:   api,
		app:   app,
		pages: pages,
		token: token,
		usrID: uID,
		fForm: NewCreateFileView(app, pages, fApi, token, uID),
	}
}

// func (c *CreateCardView) NewCreateCardForm() *tview.Form {
// 	return c.form
// }

func (c *CreateCardView) GetForm() *tview.Form {
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("authorization", *c.token))
	// Name, Login, Password, URL, Text
	// Add File - work with Name - save Card - add File, else WarningModal
	// Add Meta - work with Name - save Card - add File, else WarningModal
	c.form = tview.NewForm()
	c.form.AddInputField("Name", "", 30, nil, nil).
		AddInputField("Login", "", 30, nil, nil).
		AddInputField("Password", "", 30, nil, nil).
		AddInputField("URL", "", 30, nil, nil).
		AddInputField("Text", "", 30, nil, nil)
		// AddButton("Add File", nil).
		// AddButton("Add Meta", nil).
		// AddButton("Save", nil).
		// AddButton("Cancel", nil)
	c.form.AddButton("Save", func() {
		nameField := c.form.GetFormItemByLabel("Name").(*tview.InputField)
		loginField := c.form.GetFormItemByLabel("Login").(*tview.InputField)
		passwordField := c.form.GetFormItemByLabel("Password").(*tview.InputField)
		urlField := c.form.GetFormItemByLabel("URL").(*tview.InputField)
		textField := c.form.GetFormItemByLabel("Text").(*tview.InputField)
		name := nameField.GetText()
		login := loginField.GetText()
		password := passwordField.GetText()
		url := urlField.GetText()
		text := textField.GetText()
		_, err := c.api.CreateSecretCard(ctx, &proto.CreateSecretCardRequest{UserId: *c.usrID, Name: name, Login: login, Password: password, Url: url, Text: text})
		if err != nil {
			modal := tview.NewModal().SetText(err.Error()).AddButtons([]string{"OK"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				c.pages.RemovePage("Modal")
				c.pages.SwitchToPage("Card")
			})
			c.app.SetRoot(c.pages, true)
			c.pages.AddPage("Modal", modal, true, true)
			c.pages.SwitchToPage("Modal")
		} else {
			c.pages.SwitchToPage("Card")
		}
	})
	c.form.AddButton("Add File", func() {
		nameField := c.form.GetFormItemByLabel("Name").(*tview.InputField)
		loginField := c.form.GetFormItemByLabel("Login").(*tview.InputField)
		passwordField := c.form.GetFormItemByLabel("Password").(*tview.InputField)
		urlField := c.form.GetFormItemByLabel("URL").(*tview.InputField)
		textField := c.form.GetFormItemByLabel("Text").(*tview.InputField)
		name := nameField.GetText()
		login := loginField.GetText()
		password := passwordField.GetText()
		url := urlField.GetText()
		text := textField.GetText()
		resp, err := c.api.CreateSecretCard(ctx, &proto.CreateSecretCardRequest{UserId: *c.usrID, Name: name, Login: login, Password: password, Url: url, Text: text})
		if err != nil {
			modal := tview.NewModal().SetText(err.Error()).AddButtons([]string{"OK"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				c.pages.RemovePage("Modal")
				c.pages.SwitchToPage("Card")
			})
			c.app.SetRoot(c.pages, true)
			c.pages.AddPage("Modal", modal, true, true)
			c.pages.SwitchToPage("Modal")
		}
		if name == "" {
			modal := tview.NewModal().SetText("Name is empty").AddButtons([]string{"OK"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				c.pages.RemovePage("Modal")
				c.pages.SwitchToPage("Card")
			})
			c.app.SetRoot(c.pages, true)
			c.pages.AddPage("Modal", modal, true, true)
			c.pages.SwitchToPage("Modal")
		} else {
			c.pages.AddPage("CreateFile", c.fForm.GetForm(resp.GetId()), true, true)
			c.pages.SwitchToPage("CreateFile")
		}

	})
	return c.form
}
