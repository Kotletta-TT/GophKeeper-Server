package view

import (
	"context"
	"os"

	"github.com/Kotletta-TT/GophKeeper/proto"
	"github.com/rivo/tview"
	"google.golang.org/grpc/metadata"
)

type CreateFileView struct {
	app   *tview.Application
	pages *tview.Pages
	form  *tview.Form
	api   proto.FileSecretCardServiceClient
	token *string
	usrID *string
}

func NewCreateFileView(app *tview.Application, pages *tview.Pages, api proto.FileSecretCardServiceClient, token *string, usrID *string) *CreateFileView {
	return &CreateFileView{
		app:   app,
		pages: pages,
		api:   api,
		token: token,
		usrID: usrID,
	}
}

func (v *CreateFileView) GetForm(cardID string) *tview.Form {
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("authorization", *v.token))
	v.form = tview.NewForm()
	v.form.AddInputField("Path to file", "", 200, nil, nil)
	v.form.AddButton("Create", func() {
		path := v.form.GetFormItemByLabel("Path to file").(*tview.InputField).GetText()
		buf, err := os.ReadFile(path)
		if err != nil {
			modal := tview.NewModal().SetText(err.Error()).AddButtons([]string{"OK"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				v.pages.RemovePage("Modal")
				v.pages.SwitchToPage("Card")
			})
			v.app.SetRoot(v.pages, true)
			v.pages.AddPage("Modal", modal, true, true)
			v.pages.SwitchToPage("Modal")
			return
		}
		_, err = v.api.CreateFileSecretCard(ctx, &proto.CreateFileSecretCardRequest{
			CardId: cardID,
			File:   buf,
		})
		if err != nil {
			modal := tview.NewModal().SetText(err.Error()).AddButtons([]string{"OK"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				v.pages.RemovePage("Modal")
				v.pages.SwitchToPage("Card")
			})
			v.app.SetRoot(v.pages, true)
			v.pages.AddPage("Modal", modal, true, true)
			v.pages.SwitchToPage("Modal")
		} else {
			v.pages.SwitchToPage("Card")
		}
	})
	return v.form
}
