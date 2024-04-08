// package main

// func main() {
// 	app := tview.NewApplication()

// 	form := tview.NewForm()
// 	form.AddPasswordField("Password", "", 10, '*', nil).
// 		AddButton("Check", func() {
// 			passwordField := form.GetFormItem(0).(*tview.InputField)
// 			password := passwordField.GetText()
// 			var modal *tview.Modal
// 			if password == "password" {
// 				// Если пароль верный
// 				modal = tview.NewModal().SetText("Success").AddButtons([]string{"OK"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
// 					app.Stop()
// 				})
// 			} else {
// 				// Если пароль неверный
// 				modal = tview.NewModal().SetText("Failed").AddButtons([]string{"OK"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
// 					app.Stop()
// 				})
// 			}
// 			app.SetRoot(modal, false)
// 		}).
// 		SetButtonsAlign(tview.AlignCenter)

// 	if err := app.SetRoot(form, true).SetFocus(form).Run(); err != nil {
// 		panic(err)
// 	}
// }

package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/rivo/tview"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	app := tview.NewApplication()

	pages := tview.NewPages()
	list := tview.NewList().
		AddItem("GitHub", "", 'g', nil).
		AddItem("Google", "", 'o', nil).
		AddItem("Twitch", "", 't', nil)

	list.SetSelectedFunc(func(i int, name string, s string, t rune) {
		modal := tview.NewModal().SetText(name).AddButtons([]string{"OK"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			pages.RemovePage("LModal")
			pages.SwitchToPage("List")
		})
		app.SetRoot(pages, true)
		pages.AddPage("LModal", modal, true, true)
		pages.SwitchToPage("LModal")
	})
	form := tview.NewForm()
	form.AddInputField("Username", "", 20, nil, nil).
		AddPasswordField("Password", "", 10, '*', nil).
		AddButton("Login", func() {
			var modal *tview.Modal
			usernameField := form.GetFormItemByLabel("Username").(*tview.InputField)
			passwordField := form.GetFormItemByLabel("Password").(*tview.InputField)
			username := usernameField.GetText()
			password := passwordField.GetText()

			creds := &Credentials{
				Username: username,
				Password: password,
			}

			credsJson, _ := json.Marshal(creds)
			resp, err := http.Post("http://localhost:8000/login", "application/json", bytes.NewBuffer(credsJson))
			if err != nil {
				modal = tview.NewModal().SetText(err.Error()).AddButtons([]string{"OK"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
					pages.RemovePage("Modal")
					pages.SwitchToPage("List")
				})
				app.SetRoot(pages, true)
				pages.AddPage("Modal", modal, true, true)
				pages.SwitchToPage("Modal")
			} else {
				defer resp.Body.Close()
			}

			// Здесь вы можете обработать ответ от сервера
		}).
		SetButtonsAlign(tview.AlignCenter)
	pages.AddPage("Input", form, true, true)
	pages.AddPage("List", list, true, true)
	if err := app.SetRoot(form, true).SetFocus(form).Run(); err != nil {
		panic(err)
	}
}
