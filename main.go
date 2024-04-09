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

// package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"

// 	"github.com/rivo/tview"
// )

// type Credentials struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }

// func main() {
// 	app := tview.NewApplication()

// 	pages := tview.NewPages()
// 	list := tview.NewList().
// 		AddItem("GitHub", "", 'g', nil).
// 		AddItem("Google", "", 'o', nil).
// 		AddItem("Twitch", "", 't', nil)

// 	list.SetSelectedFunc(func(i int, name string, s string, t rune) {
// 		modal := tview.NewModal().SetText(name).AddButtons([]string{"OK"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
// 			pages.RemovePage("LModal")
// 			pages.SwitchToPage("List")
// 		})
// 		app.SetRoot(pages, true)
// 		pages.AddPage("LModal", modal, true, true)
// 		pages.SwitchToPage("LModal")
// 	})
// 	form := tview.NewForm()
// 	form.AddInputField("Username", "", 20, nil, nil).
// 		AddPasswordField("Password", "", 10, '*', nil).
// 		AddButton("Login", func() {
// 			var modal *tview.Modal
// 			usernameField := form.GetFormItemByLabel("Username").(*tview.InputField)
// 			passwordField := form.GetFormItemByLabel("Password").(*tview.InputField)
// 			username := usernameField.GetText()
// 			password := passwordField.GetText()

// 			creds := &Credentials{
// 				Username: username,
// 				Password: password,
// 			}

// 			credsJson, _ := json.Marshal(creds)
// 			resp, err := http.Post("http://localhost:8000/login", "application/json", bytes.NewBuffer(credsJson))
// 			if err != nil {
// 				modal = tview.NewModal().SetText(err.Error()).AddButtons([]string{"OK"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
// 					pages.RemovePage("Modal")
// 					pages.SwitchToPage("List")
// 				})
// 				app.SetRoot(pages, true)
// 				pages.AddPage("Modal", modal, true, true)
// 				pages.SwitchToPage("Modal")
// 			} else {
// 				defer resp.Body.Close()
// 			}

// 			// Здесь вы можете обработать ответ от сервера
// 		}).
// 		SetButtonsAlign(tview.AlignCenter)
// 	pages.AddPage("Input", form, true, true)
// 	pages.AddPage("List", list, true, true)

// 	if err := app.SetRoot(form, true).SetFocus(form).Run(); err != nil {
// 		panic(err)
// 	}
// }

package main

// import (
// 	"fmt"
// 	"os"

// 	"github.com/gdamore/tcell/v2"
// 	"github.com/rivo/tview"
// )

// func main() {
// 	// Инициализируем новое приложение tview
// 	app := tview.NewApplication()

// 	// Создаем новый примитив tview.TextView для отображения информации о файлах
// 	textView := tview.NewTextView().
// 		SetText("Перетащите файл сюда").
// 		SetTextAlign(tview.AlignCenter).
// 		SetDoneFunc(func(key tcell.Key) {
// 			// Если нажата клавиша ESC, завершаем приложение
// 			if key == tcell.KeyEscape {
// 				app.Stop()
// 			}
// 		})

// 	// Создаем примитив tview.Flex для размещения текстового представления по центру экрана
// 	flex := tview.NewFlex().
// 		AddItem(nil, 0, 1, false).
// 		AddItem(textView, 0, 1, true).
// 		AddItem(nil, 0, 1, false)

// 	// Устанавливаем функцию для обработки событий мыши
// 	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
// 		if event.Key() == tcell.KeyRune {
// 			// Если нажата клавиша и это не специальная клавиша (например, Esc или Enter), выводим информацию о событии
// 			rune := event.Rune()
// 			if rune == 'D' || rune == 'd' {
// 				if event.Modifiers()&tcell.ModShift > 0 {
// 					app.SetRoot(textView, true)
// 					textView.SetText("Файл был перетащен сюда")
// 				}
// 			}
// 		}
// 		// Возвращаем nil, чтобы событие не обрабатывалось другими обработчиками
// 		return nil
// 	})

// 	// Запускаем приложение
// 	if err := app.SetRoot(flex, true).Run(); err != nil {
// 		fmt.Fprintf(os.Stderr, "%v\n", err)
// 		os.Exit(1)
// 	}
// }
